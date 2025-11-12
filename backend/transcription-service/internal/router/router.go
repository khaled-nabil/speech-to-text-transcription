package router

import (
	"transcription-service/controller/transcribercontroller"

	"github.com/gin-gonic/gin"

	"transcription-service/controller/healthcontroller"
)

type Router struct {
	Gin   *gin.Engine
	hctrl *healthcontroller.Controller
	tctrl *transcribercontroller.Controller
}

func New(
	gin *gin.Engine,
	hctrl *healthcontroller.Controller,
	tctrl *transcribercontroller.Controller,
) *Router {
	return &Router{gin, hctrl, tctrl}
}

func (r *Router) Route() {
	r.Gin.Group("/api/v1/health").
		GET("", r.hctrl.Health)

	r.Gin.Group("/api/v1/transcriber").
		GET(":filename", r.tctrl.GetAudio).
		POST("", r.tctrl.UploadAudio)
}
