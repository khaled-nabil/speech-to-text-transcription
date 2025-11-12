package transcription

import (
	"mime/multipart"
)

type UseCase interface {
	UploadAudio(userID string, fileHeader *multipart.FileHeader) (string, error)
	GetAudio(userID string, fileName string) ([]byte, string, error)
}
