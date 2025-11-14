package transcribercontroller

import "time"

type (
	STATUS string

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
