package transcription

import (
	"mime/multipart"
	"transcription-service/domain/transcriptionentity"
)

type UseCase interface {
	UploadAudio(userID string, fileHeader *multipart.FileHeader) (id string, err error)
	GetAudio(userID string, fileName string) (file []byte, contentType string, err error)
	GetAllUserTranscriptions(userID string) ([]*transcriptionentity.Transcription, error)
	GetTranscriptionByID(id string) (*transcriptionentity.Transcription, error)
}
