package transcription

import (
	"bytes"
	"errors"
	"fmt"
	"mime/multipart"
	"net/textproto"
	"testing"
	"time"
	persistancemocks "transcription-service/domain/persistance/mocks"
	"transcription-service/domain/transcriber"
	transcribermocks "transcription-service/domain/transcriber/mocks"
	"transcription-service/domain/transcriptionentity"
	"transcription-service/internal/config"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func createFileHeader(filename string, content []byte, mimeType string) *multipart.FileHeader {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, filename))
	if mimeType != "" {
		h.Set("Content-Type", mimeType)
	}

	part, _ := writer.CreatePart(h)
	part.Write(content)
	writer.Close()

	reader := multipart.NewReader(body, writer.Boundary())
	form, _ := reader.ReadForm(10 << 20)

	return form.File["file"][0]
}

func TestUploadAudio(t *testing.T) {
	t.Run("When file size exceeds allowed, fail with correct message", func(t *testing.T) {
		mockStorage := persistancemocks.NewMockStorage(t)
		mockRepo := persistancemocks.NewMockTranscriptionRepository(t)
		mockTranscriber := transcribermocks.NewMockTranscriber(t)

		cfg := &config.Config{
			Storage: config.StorageConfig{
				MaxFileSize: 5,
				AllowedExt:  []string{"mp3", "wav"},
			},
		}

		uc := New(mockStorage, mockRepo, mockTranscriber, cfg)

		largeContent := make([]byte, 6*1024*1024)
		fileHeader := createFileHeader("test.mp3", largeContent, "audio/mpeg")

		id, err := uc.UploadAudio("user123", fileHeader)

		assert.Error(t, err)
		assert.Empty(t, id)
		assert.Contains(t, err.Error(), "file too large")
	})

	t.Run("When file format is not supported, fail with correct message", func(t *testing.T) {
		mockStorage := persistancemocks.NewMockStorage(t)
		mockRepo := persistancemocks.NewMockTranscriptionRepository(t)
		mockTranscriber := transcribermocks.NewMockTranscriber(t)

		cfg := &config.Config{
			Storage: config.StorageConfig{
				MaxFileSize: 10,
				AllowedExt:  []string{"mp3", "wav"},
			},
		}

		uc := New(mockStorage, mockRepo, mockTranscriber, cfg)

		content := []byte("fake audio content")
		fileHeader := createFileHeader("test.xyz", content, "audio/xyz")

		id, err := uc.UploadAudio("user123", fileHeader)

		assert.Error(t, err)
		assert.Empty(t, id)
		assert.Contains(t, err.Error(), "not allowed")
	})

	t.Run("When userID is missing, return error", func(t *testing.T) {
		mockStorage := persistancemocks.NewMockStorage(t)
		mockRepo := persistancemocks.NewMockTranscriptionRepository(t)
		mockTranscriber := transcribermocks.NewMockTranscriber(t)

		cfg := &config.Config{
			Storage: config.StorageConfig{
				MaxFileSize: 10,
				AllowedExt:  []string{"mp3"},
			},
		}

		uc := New(mockStorage, mockRepo, mockTranscriber, cfg)

		content := []byte("test")
		fileHeader := createFileHeader("test.mp3", content, "audio/mpeg")

		id, err := uc.UploadAudio("", fileHeader)

		assert.Error(t, err)
		assert.Empty(t, id)
		assert.Contains(t, err.Error(), "missing userID")
	})

	t.Run("When file is valid and async processing succeeds, upload and return ID", func(t *testing.T) {
		mockStorage := persistancemocks.NewMockStorage(t)
		mockRepo := persistancemocks.NewMockTranscriptionRepository(t)
		mockTranscriber := transcribermocks.NewMockTranscriber(t)

		cfg := &config.Config{
			Storage: config.StorageConfig{
				MaxFileSize: 10,
				AllowedExt:  []string{"mp3", "wav"},
			},
		}

		uc := New(mockStorage, mockRepo, mockTranscriber, cfg)

		content := []byte("fake audio content")
		fileHeader := createFileHeader("test.mp3", content, "audio/mpeg")

		mockStorage.EXPECT().StoreFile(mock.AnythingOfType("string"), content).Return(nil)
		mockRepo.EXPECT().Save(mock.AnythingOfType("*transcriptionentity.Transcription")).Return(nil)

		mockTranscriber.EXPECT().Transcribe(mock.Anything).Return(
			&transcriber.TranscriptionResponse{Text: "test transcription"}, nil,
		).Maybe()
		mockRepo.EXPECT().GetByID(mock.AnythingOfType("string")).Return(
			&transcriptionentity.Transcription{ID: "test"}, nil,
		).Maybe()
		mockRepo.EXPECT().UpdateWithTranscription(mock.Anything).Return(nil).Maybe()

		id, err := uc.UploadAudio("user123", fileHeader)

		assert.NoError(t, err)
		assert.NotEmpty(t, id)

		time.Sleep(10 * time.Millisecond) // Allow goroutine to complete
	})

	t.Run("When repository save fails, return error", func(t *testing.T) {
		mockStorage := persistancemocks.NewMockStorage(t)
		mockRepo := persistancemocks.NewMockTranscriptionRepository(t)
		mockTranscriber := transcribermocks.NewMockTranscriber(t)

		cfg := &config.Config{
			Storage: config.StorageConfig{
				MaxFileSize: 10,
				AllowedExt:  []string{"mp3"},
			},
		}

		uc := New(mockStorage, mockRepo, mockTranscriber, cfg)

		content := []byte("test")
		fileHeader := createFileHeader("test.mp3", content, "audio/mpeg")

		mockStorage.EXPECT().StoreFile(mock.AnythingOfType("string"), content).Return(nil)
		mockRepo.EXPECT().Save(mock.AnythingOfType("*transcriptionentity.Transcription")).Return(errors.New("db error"))

		id, err := uc.UploadAudio("user123", fileHeader)

		assert.Error(t, err)
		assert.Empty(t, id)
		assert.Contains(t, err.Error(), "failed to save transcription")
	})
}

func TestGetTranscriptionByID(t *testing.T) {
	t.Run("When transcription exists, return transcription with data", func(t *testing.T) {
		mockStorage := persistancemocks.NewMockStorage(t)
		mockRepo := persistancemocks.NewMockTranscriptionRepository(t)
		mockTranscriber := transcribermocks.NewMockTranscriber(t)

		cfg := &config.Config{}
		uc := New(mockStorage, mockRepo, mockTranscriber, cfg)

		expectedTranscription := &transcriptionentity.Transcription{
			ID:             "trans123",
			UserID:         "user123",
			FileName:       "test.mp3",
			UploadDate:     time.Now(),
			TranscriptText: "Hello world",
			Status:         transcriptionentity.StatusSuccess,
		}

		mockRepo.EXPECT().GetByID("trans123").Return(expectedTranscription, nil)

		result, err := uc.GetTranscriptionByID("trans123")

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedTranscription.ID, result.ID)
		assert.Equal(t, expectedTranscription.UserID, result.UserID)
		assert.Equal(t, expectedTranscription.TranscriptText, result.TranscriptText)
	})

	t.Run("When transcription ID is empty, return error", func(t *testing.T) {
		mockStorage := persistancemocks.NewMockStorage(t)
		mockRepo := persistancemocks.NewMockTranscriptionRepository(t)
		mockTranscriber := transcribermocks.NewMockTranscriber(t)

		cfg := &config.Config{}
		uc := New(mockStorage, mockRepo, mockTranscriber, cfg)

		result, err := uc.GetTranscriptionByID("")

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "missing transcription ID")
	})

	t.Run("When transcription does not exist, return error", func(t *testing.T) {
		mockStorage := persistancemocks.NewMockStorage(t)
		mockRepo := persistancemocks.NewMockTranscriptionRepository(t)
		mockTranscriber := transcribermocks.NewMockTranscriber(t)

		cfg := &config.Config{}
		uc := New(mockStorage, mockRepo, mockTranscriber, cfg)

		mockRepo.EXPECT().GetByID("nonexistent").Return(nil, errors.New("not found"))

		result, err := uc.GetTranscriptionByID("nonexistent")

		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestGetAllUserTranscriptions(t *testing.T) {
	t.Run("When user has transcriptions, return all transcriptions", func(t *testing.T) {
		mockStorage := persistancemocks.NewMockStorage(t)
		mockRepo := persistancemocks.NewMockTranscriptionRepository(t)
		mockTranscriber := transcribermocks.NewMockTranscriber(t)

		cfg := &config.Config{}
		uc := New(mockStorage, mockRepo, mockTranscriber, cfg)

		expectedTranscriptions := []*transcriptionentity.Transcription{
			{ID: "trans1", UserID: "user123", Status: transcriptionentity.StatusSuccess},
			{ID: "trans2", UserID: "user123", Status: transcriptionentity.StatusPending},
		}

		mockRepo.EXPECT().GetAllByUserID("user123").Return(expectedTranscriptions, nil)

		result, err := uc.GetAllUserTranscriptions("user123")

		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, "trans1", result[0].ID)
		assert.Equal(t, "trans2", result[1].ID)
	})

	t.Run("When userID is empty, return error", func(t *testing.T) {
		mockStorage := persistancemocks.NewMockStorage(t)
		mockRepo := persistancemocks.NewMockTranscriptionRepository(t)
		mockTranscriber := transcribermocks.NewMockTranscriber(t)

		cfg := &config.Config{}
		uc := New(mockStorage, mockRepo, mockTranscriber, cfg)

		result, err := uc.GetAllUserTranscriptions("")

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "missing userID")
	})
}

func TestGetAudio(t *testing.T) {
	t.Run("When audio exists, return audio data and MIME type", func(t *testing.T) {
		mockStorage := persistancemocks.NewMockStorage(t)
		mockRepo := persistancemocks.NewMockTranscriptionRepository(t)
		mockTranscriber := transcribermocks.NewMockTranscriber(t)

		cfg := &config.Config{}
		uc := New(mockStorage, mockRepo, mockTranscriber, cfg)

		expectedData := []byte("audio content")
		mockStorage.EXPECT().GetFile("user123/test.mp3").Return(expectedData, nil)

		data, mimeType, err := uc.GetAudio("user123", "test.mp3")

		assert.NoError(t, err)
		assert.Equal(t, expectedData, data)
		assert.Equal(t, "audio/mpeg", mimeType)
	})

	t.Run("When userID or fileName is missing, return error", func(t *testing.T) {
		mockStorage := persistancemocks.NewMockStorage(t)
		mockRepo := persistancemocks.NewMockTranscriptionRepository(t)
		mockTranscriber := transcribermocks.NewMockTranscriber(t)

		cfg := &config.Config{}
		uc := New(mockStorage, mockRepo, mockTranscriber, cfg)

		data, mimeType, err := uc.GetAudio("", "test.mp3")

		assert.Error(t, err)
		assert.Nil(t, data)
		assert.Empty(t, mimeType)
		assert.Contains(t, err.Error(), "missing identifiers")
	})

	t.Run("When storage fails, return error", func(t *testing.T) {
		mockStorage := persistancemocks.NewMockStorage(t)
		mockRepo := persistancemocks.NewMockTranscriptionRepository(t)
		mockTranscriber := transcribermocks.NewMockTranscriber(t)

		cfg := &config.Config{}
		uc := New(mockStorage, mockRepo, mockTranscriber, cfg)

		mockStorage.EXPECT().GetFile("user123/test.mp3").Return(nil, errors.New("storage error"))

		data, mimeType, err := uc.GetAudio("user123", "test.mp3")

		assert.Error(t, err)
		assert.Nil(t, data)
		assert.Empty(t, mimeType)
	})
}
