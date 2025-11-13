package persistance

import (
	"context"
	"transcription-service/domain/transcriptionentity"
)

type (
	Storage interface {
		StoreFile(path string, data []byte) error
		GetFile(path string) ([]byte, error)
		DeleteFile(path string) error
	}

	TranscriptionRepository interface {
		Save(ctx context.Context, transcription *transcriptionentity.Transcription) error
		GetByID(ctx context.Context, id string) (*transcriptionentity.Transcription, error)
		GetAllByUserID(ctx context.Context, userID string) ([]*transcriptionentity.Transcription, error)
		Close() error
	}
)
