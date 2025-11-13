package transcriber

type (
	Transcriber interface {
		Transcribe(req TranscriptionRequest) (*TranscriptionResponse, error)
	}

	TranscriptionRequest struct {
		Temperature float64
		FileName    string
		Language    string
		Prompt      string
		AudioFile   []byte
	}

	TranscriptionResponse struct {
		Text string `json:"text"`
	}

	ErrorResponse struct {
		Detail struct {
			Message string `json:"msg"`
			Type    string `json:"type"`
		} `json:"detail"`
	}
)
