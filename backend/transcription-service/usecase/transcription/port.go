package transcription

import (
	"mime/multipart"
)

type UseCase interface {
	GetTranscription(userID string, fileHeader *multipart.FileHeader) (filename string, transcription string, err error)
	GetAudio(userID string, fileName string) (file []byte, contentType string, err error)
}
