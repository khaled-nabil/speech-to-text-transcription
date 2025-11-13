package transcription

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"path"
	"slices"
	"strings"
	"time"
	"transcription-service/domain/persistance"
	"transcription-service/domain/transcriber"
	"transcription-service/domain/transcriptionentity"
	"transcription-service/internal/config"

	"github.com/google/uuid"
)

type useCase struct {
	storage     persistance.Storage
	repo        persistance.TranscriptionRepository
	transcriber transcriber.Transcriber
	config      *config.StorageConfig
}

func New(storage persistance.Storage, repo persistance.TranscriptionRepository, transcriber transcriber.Transcriber, config *config.Config) UseCase {
	return &useCase{
		storage:     storage,
		repo:        repo,
		transcriber: transcriber,
		config:      &config.Storage,
	}
}

func (u *useCase) GetTranscription(userID string, fileHeader *multipart.FileHeader) (string, string, error) {
	if userID == "" {
		return "", "", errors.New("missing userID")
	}
	if fileHeader == nil {
		return "", "", errors.New("missing file")
	}

	maxFileSize := maxFileSizeFromMBytes(u.config.MaxFileSize)

	if fileHeader.Size > maxFileSize {
		return "", "", fmt.Errorf("file too large (>%dMB)", maxFileSize)
	}

	var fileExtension string

	if mimeType := fileHeader.Header.Get("Content-Type"); mimeType != "" {
		fileExtension, _ = getExtensionFromMIME(mimeType)
	}
	if fileExtension == "" {
		fileExtension = strings.TrimPrefix(path.Ext(fileHeader.Filename), ".")
	}

	if (slices.Contains(u.config.AllowedExt, fileExtension)) == false {
		return "", "", errors.New(fmt.Sprintf("file type %s not allowed", fileExtension))
	}

	f, err := fileHeader.Open()
	if err != nil {
		return "", "", err
	}
	defer func(f multipart.File) {
		err = f.Close()
		if err != nil {
			log.Printf("Error closing file: %v", err)
		}
	}(f)

	data, err := io.ReadAll(io.LimitReader(f, maxFileSize+1))
	if err != nil {
		return "", "", err
	}
	if int64(len(data)) > maxFileSize {
		return "", "", errors.New(fmt.Sprintf("file too large (>%dMB)", maxFileSize))
	}

	fileName := fmt.Sprintf("%s.%s", uuid.New().String(), fileExtension)
	objectPath := fmt.Sprintf("%s/%s", userID, fileName)

	if err = u.storage.StoreFile(objectPath, data); err != nil {
		return "", "", err
	}

	resp, err := u.transcriber.Transcribe(transcriber.TranscriptionRequest{
		FileName:  fileName,
		AudioFile: data,
	})
	if err != nil {
		return "", "", err
	}

	transcriptionEntity := &transcriptionentity.Transcription{
		ID:             uuid.New().String(),
		UserID:         userID,
		FileName:       fileName,
		UploadDate:     time.Now(),
		TranscriptText: resp.Text,
	}
	if err = u.repo.Save(context.Background(), transcriptionEntity); err != nil {
		return "", "", fmt.Errorf("failed to save transcription: %w", err)
	}

	return fileName, resp.Text, nil
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

	mime, err := getMIMEFromExtension(strings.ToLower(path.Ext(fileName)))
	if err != nil {
		return nil, "", errors.New("unable to determine file MIME type")
	}

	return data, mime, nil
}

func (u *useCase) GetAllUserTranscriptions(userID string) ([]*transcriptionentity.Transcription, error) {
	if userID == "" {
		return nil, errors.New("missing userID")
	}

	return u.repo.GetAllByUserID(context.Background(), userID)
}

func (u *useCase) GetTranscriptionByID(id string) (*transcriptionentity.Transcription, error) {
	if id == "" {
		return nil, errors.New("missing transcription ID")
	}

	return u.repo.GetByID(context.Background(), id)
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
