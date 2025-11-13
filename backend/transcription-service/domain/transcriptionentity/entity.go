package transcriptionentity

import "time"

type Transcription struct {
	ID             string
	UserID         string
	FileName       string
	UploadDate     time.Time
	TranscriptText string
}
