package transcribercontroller

import "time"

type (
	STATUS string

	TranscriptionResponseDTO struct {
		ID     string `json:"id"`
		Status STATUS `json:"status"`
	}

	TranscriptionItemDTO struct {
		ID             string    `json:"id"`
		UploadDate     time.Time `json:"uploadDate"`
		TranscriptText string    `json:"transcriptText"`
		Status         STATUS    `json:"status"`
	}
)

const (
	PENDING STATUS = "PENDING"
	SUCCESS STATUS = "SUCCESS"
	ERROR   STATUS = "ERROR"
)
