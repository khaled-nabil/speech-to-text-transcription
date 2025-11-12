package transcribercontroller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func New() *Controller {
	return &Controller{}
}

func (*Controller) GetFile(c *gin.Context) {
	// return current timestamp
	c.JSON(http.StatusCreated, gin.H{"status": "ok", "timestamp": time.Now().Unix()})
}
