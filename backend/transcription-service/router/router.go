package router

import (
	"github.com/gin-gonic/gin"

	"transcription-service/controller/healthcontroller"
)

type Router struct {
	Gin   *gin.Engine
	hctrl *healthcontroller.Controller
}

func New(
	gin *gin.Engine,
	hctrl *healthcontroller.Controller,
) *Router {
	return &Router{gin, hctrl}
}

func (r *Router) Route() {
	r.Gin.Group("/api/v1").
		GET("/health", r.hctrl.Health)
}
