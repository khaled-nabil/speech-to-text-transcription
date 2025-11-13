package transcribercontroller

import "time"

type (
	STATUS string

	TranscriptionResponseDTO struct {
		Text     string `json:"text"`
		FileName string `json:"fileName"`
		Status   STATUS `json:"status"`
	}

	TranscriptionItemDTO struct {
		ID             string    `json:"id"`
		UploadDate     time.Time `json:"uploadDate"`
		TranscriptText string    `json:"transcriptText"`
	}
)

const (
	PENDING STATUS = "PENDING"
	SUCCESS STATUS = "SUCCESS"
	ERROR   STATUS = "ERROR"
)
