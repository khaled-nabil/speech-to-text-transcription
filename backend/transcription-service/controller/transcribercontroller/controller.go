package transcribercontroller

import (
	"net/http"

	"transcription-service/usecase/transcription"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	useCase transcription.UseCase
}

func New(u transcription.UseCase) *Controller {
	return &Controller{useCase: u}
}

func (ctr *Controller) UploadAudio(c *gin.Context) {
	userID := c.GetHeader("X-User-ID")
	fileHeader, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file missing"})
		return
	}

	t, err := ctr.useCase.UploadAudio(userID, fileHeader)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	signedURL, err := ctr.useCase.GetPresignedURL(userID, t.FileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, TranscriptionToDTO(t, signedURL))
}

func (ctr *Controller) GetAudio(c *gin.Context) {
	fileName := c.Param("filename")
	userID := c.GetHeader("X-User-ID")

	data, contentType, err := ctr.useCase.GetAudio(userID, fileName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, contentType, data)
}

func (ctr *Controller) GetAllUserTranscriptions(c *gin.Context) {
	userID := c.GetHeader("X-User-ID")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing user ID"})
		return
	}

	transcriptions, err := ctr.useCase.GetAllUserTranscriptions(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []TranscriptionItemDTO
	for _, t := range transcriptions {
		response = append(response, TranscriptionItemDTO{
			ID:             t.ID,
			UploadDate:     t.UploadDate,
			TranscriptText: t.TranscriptText,
			Status:         STATUS(t.Status),
		})
	}

	c.JSON(http.StatusOK, response)
}

func (ctr *Controller) GetTranscriptionByID(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetHeader("X-User-ID")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing transcription ID"})
		return
	}

	t, err := ctr.useCase.GetTranscriptionByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	signedURL, err := ctr.useCase.GetPresignedURL(userID, t.FileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, TranscriptionToDTO(t, signedURL))
}
