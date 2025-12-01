package persistance

import (
	"transcription-service/domain/transcriptionentity"
)

type (
	Storage interface {
		StoreFile(path string, data []byte) error
		GetFile(path string) ([]byte, error)
		GetPresignedURL(path string) (string, error)
		DeleteFile(path string) error
	}

	TranscriptionRepository interface {
		Save(transcription *transcriptionentity.Transcription) error
		UpdateWithTranscription(transcription *transcriptionentity.Transcription) error
		GetByID(id string) (*transcriptionentity.Transcription, error)
		GetAllByUserID(userID string) ([]*transcriptionentity.Transcription, error)
		Close() error
	}
)
