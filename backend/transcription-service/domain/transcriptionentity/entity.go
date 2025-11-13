package transcriptionentity

import "time"

type Status string

const (
	StatusPending Status = "PENDING"
	StatusSuccess Status = "SUCCESS"
	StatusError   Status = "ERROR"
)

type Transcription struct {
	ID             string
	UserID         string
	FileName       string
	UploadDate     time.Time
	TranscriptText string
	Status         Status
}
