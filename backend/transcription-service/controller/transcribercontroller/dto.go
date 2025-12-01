package transcribercontroller

import (
	"time"
	"transcription-service/domain/transcriptionentity"
)

type (
	STATUS string

	TranscriptionItemDTO struct {
		ID             string    `json:"id"`
		UploadDate     time.Time `json:"uploadDate"`
		MediaURL       string    `json:"mediaUrl"`
		TranscriptText string    `json:"transcriptText,omitempty"`
		Status         STATUS    `json:"status"`
	}
)

const (
	PENDING STATUS = "PENDING"
	SUCCESS STATUS = "SUCCESS"
	ERROR   STATUS = "ERROR"
)

func TranscriptionToDTO(t *transcriptionentity.Transcription, url string) *TranscriptionItemDTO {
	return &TranscriptionItemDTO{
		ID:             t.ID,
		UploadDate:     t.UploadDate,
		MediaURL:       url,
		TranscriptText: t.TranscriptText,
		Status:         STATUS(t.Status),
	}
}
