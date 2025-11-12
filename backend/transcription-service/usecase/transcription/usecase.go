package transcription

import (
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"path"
	"slices"
	"strings"
	"transcription-service/internal/config"

	"transcription-service/pkg/minio"

	"github.com/google/uuid"
)

type useCase struct {
	storage *minio.Storage
	config  *config.StorageConfig
}

func New(storage *minio.Storage, config *config.Config) UseCase {
	return &useCase{storage: storage, config: &config.Storage}
}

func (u *useCase) UploadAudio(userID string, fileHeader *multipart.FileHeader) (string, error) {
	if userID == "" {
		return "", errors.New("missing userID")
	}
	if fileHeader == nil {
		return "", errors.New("missing file")
	}
	if fileHeader.Size > maxFileSizeFromMBytes(u.config.MaxFileSize) {
		return "", errors.New("file too large (>5MB)")
	}

	var fileExtension string

	if mimeType := fileHeader.Header.Get("Content-Type"); mimeType != "" {
		fileExtension, _ = getExtensionFromMIME(mimeType)
	}
	if fileExtension == "" {
		fileExtension = strings.TrimPrefix(path.Ext(fileHeader.Filename), ".")
	}

	if (slices.Contains(u.config.AllowedExt, fileExtension)) == false {
		return "", errors.New(fmt.Sprintf("file type %s not allowed", fileExtension))
	}

	f, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer func(f multipart.File) {
		err = f.Close()
		if err != nil {
			log.Printf("Error closing file: %v", err)
		}
	}(f)

	data, err := io.ReadAll(io.LimitReader(f, maxFileSizeFromMBytes(u.config.MaxFileSize)+1))
	if err != nil {
		return "", err
	}
	if int64(len(data)) > maxFileSizeFromMBytes(u.config.MaxFileSize) {
		return "", errors.New(fmt.Sprintf("file too large (>%dMB)", maxFileSizeFromMBytes(u.config.MaxFileSize)))
	}

	fileName := uuid.New().String()
	objectPath := fmt.Sprintf("%s/%s.%s", userID, fileName, fileExtension)

	if err = u.storage.StoreFile(objectPath, data); err != nil {
		return "", err
	}
	return fileName, nil
}

func (u *useCase) GetAudio(userID string, fileName string) ([]byte, string, error) {
	if userID == "" || fileName == "" {
		return nil, "", errors.New("missing identifiers")
	}

	objectPath := fmt.Sprintf("%s/%s", userID, fileName)
	data, err := u.storage.GetFile(objectPath)
	if err != nil {
		return nil, "", err
	}

	ext, err := getMIMEFromExtension(strings.ToLower(path.Ext(fileName)))
	if err != nil {
		return nil, "", errors.New("unable to determine file MIME type")
	}

	return data, ext, nil
}

func maxFileSizeFromMBytes(mbytes int) int64 {
	return int64(mbytes) * 1024 * 1024
}

func getExtensionFromMIME(m string) (string, error) {
	switch m {
	case "audio/mpeg":
		return "mp3", nil
	case "audio/wav", "audio/x-wav":
		return "wav", nil
	case "audio/x-m4a", "audio/mp4":
		return "m4a", nil
	case "audio/flac":
		return "flac", nil
	case "audio/ogg":
		return "ogg", nil
	default:
		return "", fmt.Errorf("unsupported MIME type: %s", m)
	}
}

func getMIMEFromExtension(name string) (string, error) {
	ext := strings.ToLower(path.Ext(name))
	switch ext {
	case "mp3":
		return "audio/mpeg", nil
	case "wav":
		return "audio/wav", nil
	case "m4a":
		return "audio/x-m4a", nil
	case "flac":
		return "audio/flac", nil
	case "ogg":
		return "audio/ogg", nil
	default:
		return "", fmt.Errorf("unsupported file extension: %s", ext)
	}
}
