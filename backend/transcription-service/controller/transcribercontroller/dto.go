package transcribercontroller

type (
	STATUS string

	TranscriptionResponseDTO struct {
		Text     string `json:"text"`
		FileName string `json:"fileName"`
		Status   STATUS `json:"status"`
	}
)

const (
	PENDING STATUS = "PENDING"
	SUCCESS STATUS = "SUCCESS"
	ERROR   STATUS = "ERROR"
)
