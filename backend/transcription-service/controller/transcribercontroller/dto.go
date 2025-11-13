package transcribercontroller

type TranscriptionResponseDTO struct {
	Text     string `json:"text"`
	FileName string `json:"fileName"`
}
