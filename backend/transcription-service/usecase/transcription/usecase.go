package transcription

import (
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

func (u *useCase) UploadAudio(userID string, fileHeader *multipart.FileHeader) (*transcriptionentity.Transcription, error) {
	if userID == "" {
		return nil, fmt.Errorf("missing userID")
	}
	if fileHeader == nil {
		return nil, fmt.Errorf("missing file")
	}

	maxFileSize := maxFileSizeFromMBytes(u.config.MaxFileSize)

	if fileHeader.Size > maxFileSize {
		return nil, fmt.Errorf("file too large (>%dMB)", maxFileSize)
	}

	var fileExtension string

	if mimeType := fileHeader.Header.Get("Content-Type"); mimeType != "" {
		fileExtension, _ = getExtensionFromMIME(mimeType)
	}
	if fileExtension == "" {
		fileExtension = strings.TrimPrefix(path.Ext(fileHeader.Filename), ".")
	}

	if !(slices.Contains(u.config.AllowedExt, fileExtension)) {
		return nil, fmt.Errorf("file type %s not allowed", fileExtension)
	}

	f, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer func(f multipart.File) {
		err = f.Close()
		if err != nil {
			log.Printf("Error closing file: %v", err)
		}
	}(f)

	data, err := io.ReadAll(io.LimitReader(f, maxFileSize+1))
	if err != nil {
		return nil, err
	}
	if int64(len(data)) > maxFileSize {
		return nil, fmt.Errorf("file too large (>%dMB)", maxFileSize)
	}

	fileName := fmt.Sprintf("%s.%s", uuid.New().String(), fileExtension)
	objectPath := fmt.Sprintf("%s/%s", userID, fileName)

	if err = u.storage.StoreFile(objectPath, data); err != nil {
		return nil, err
	}

	transcriptionEntity := &transcriptionentity.Transcription{
		ID:             uuid.New().String(),
		UserID:         userID,
		FileName:       fileName,
		UploadDate:     time.Now(),
		TranscriptText: "",
		Status:         transcriptionentity.StatusPending,
	}
	if err = u.repo.Save(transcriptionEntity); err != nil {
		return nil, fmt.Errorf("failed to save transcription: %w", err)
	}

	go u.processTranscription(transcriptionEntity.ID, fileName, data)

	return transcriptionEntity, nil
}

func (u *useCase) processTranscription(id string, fileName string, data []byte) {
	resp, err := u.transcriber.Transcribe(transcriber.TranscriptionRequest{
		FileName:  fileName,
		AudioFile: data,
	})

	transcription, getErr := u.repo.GetByID(id)
	if getErr != nil {
		log.Printf("Error getting transcription %s: %v", id, getErr)
		return
	}

	if err != nil {
		log.Printf("Error transcribing audio for %s: %v", id, err)
		transcription.Status = transcriptionentity.StatusError
		transcription.TranscriptText = fmt.Sprintf("Transcription failed: %v", err)
	} else {
		transcription.Status = transcriptionentity.StatusSuccess
		transcription.TranscriptText = resp.Text
	}

	if updateErr := u.repo.UpdateWithTranscription(transcription); updateErr != nil {
		log.Printf("Error updating transcription %s: %v", id, updateErr)
	}
}

func (u *useCase) GetAudio(userID string, fileName string) ([]byte, string, error) {
	if userID == "" || fileName == "" {
		return nil, "", fmt.Errorf("missing identifiers")
	}

	objectPath := fmt.Sprintf("%s/%s", userID, fileName)
	data, err := u.storage.GetFile(objectPath)
	if err != nil {
		return nil, "", err
	}

	mime, err := getMIMEFromExtension(strings.ToLower(path.Ext(fileName)[1:]))
	if err != nil {
		return nil, "", fmt.Errorf("unable to determine file MIME type")
	}

	return data, mime, nil
}

func (u *useCase) GetAllUserTranscriptions(userID string) ([]*transcriptionentity.Transcription, error) {
	if userID == "" {
		return nil, fmt.Errorf("missing userID")
	}

	return u.repo.GetAllByUserID(userID)
}

func (u *useCase) GetTranscriptionByID(id string) (*transcriptionentity.Transcription, error) {
	if id == "" {
		return nil, fmt.Errorf("missing transcription ID")
	}

	return u.repo.GetByID(id)
}

func maxFileSizeFromMBytes(mbytes int) int64 {
	return int64(mbytes) * 1024 * 1024
}

func getExtensionFromMIME(m string) (string, error) {
	switch m {
	case "audio/webm":
		return "webm", nil
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

func getMIMEFromExtension(ext string) (string, error) {
	switch ext {
	case "webm":
		return "audio/webm", nil
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
