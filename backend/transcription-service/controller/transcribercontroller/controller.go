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

	f, t, err := ctr.useCase.GetTranscription(userID, fileHeader)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, TranscriptionResponseDTO{
		Text:     t,
		FileName: f,
		Status:   SUCCESS,
	})
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
