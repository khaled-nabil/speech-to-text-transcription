package faster_whisper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"time"
	"transcription-service/domain/transcriber"
	"transcription-service/internal/config"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

const (
	ApiPrefix             = "/v1"
	DownloadModelEndpoint = "/models"
	TranscriptionEndpoint = "/audio/transcriptions"
)

func New(c *config.Config) (*Client, error) {
	baseURL := fmt.Sprintf("%s:%s", c.FasterWhisper.Endpoint, c.FasterWhisper.Port)
	httpC := &http.Client{
		Timeout: 5 * time.Minute,
	}

	// download model into the faster-whisper server
	log.Printf("Downloading model %s, this can take a little while..", c.FasterWhisper.Model)
	_, err := httpC.Post(fmt.Sprint(baseURL, ApiPrefix, DownloadModelEndpoint, c.FasterWhisper.Model), "application/json", nil)
	if err != nil {
		// break since the model is required for the transcription service to work
		panic(fmt.Errorf("failed to download model: %w", err))
	}
	log.Println("Model downloaded successfully - ready to transcribe!")

	return &Client{
		baseURL:    baseURL,
		httpClient: httpC,
	}, nil
}

// Transcribe sends an audio file to faster-whisper and returns the transcription
func (c *Client) Transcribe(req transcriber.TranscriptionRequest) (*transcriber.TranscriptionResponse, error) {
	body, contentType, err := getRequestBody(req)

	resp, err := c.httpClient.Post(fmt.Sprint(c.baseURL, ApiPrefix, TranscriptionEndpoint), contentType, body)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Printf("Error closing response body: %v", err)
		}
	}(resp.Body)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var errResp transcriber.ErrorResponse
		if err = json.Unmarshal(respBody, &errResp); err != nil {
			return nil, fmt.Errorf("transcription failed with status %d: %s", resp.StatusCode, string(respBody))
		}
		return nil, fmt.Errorf("transcription failed: %s", errResp.Detail.Message)
	}

	var transcriptionResp transcriber.TranscriptionResponse
	if err = json.Unmarshal(respBody, &transcriptionResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &transcriptionResp, nil
}

func getRequestBody(req transcriber.TranscriptionRequest) (*bytes.Buffer, string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("audio_file", req.FileName)
	if err != nil {
		return nil, "", fmt.Errorf("failed to create form file: %w", err)
	}
	if _, err = io.Copy(part, bytes.NewReader(req.AudioFile)); err != nil {
		return nil, "", fmt.Errorf("failed to write audio file: %w", err)
	}

	if req.Language != "" {
		if err = writer.WriteField("language", req.Language); err != nil {
			return nil, "", fmt.Errorf("failed to write language field: %w", err)
		}
	}

	if req.Temperature > 0 {
		if err = writer.WriteField("temperature", fmt.Sprintf("%.2f", req.Temperature)); err != nil {
			return nil, "", fmt.Errorf("failed to write temperature field: %w", err)
		}
	}

	if req.Prompt != "" {
		if err = writer.WriteField("prompt", req.Prompt); err != nil {
			return nil, "", fmt.Errorf("failed to write prompt: %w", err)
		}
	}

	if err = writer.Close(); err != nil {
		return nil, "", fmt.Errorf("failed to close writer: %w", err)
	}

	return body, writer.FormDataContentType(), nil
}
