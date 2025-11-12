package server

import (
	"fmt"
	"transcription-service/router"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	router *router.Router
}

func New(e *gin.Engine, r *router.Router) *Server {
	return &Server{e, r}
}

func (s *Server) Start() error {

	s.engine.Use(gin.Logger())
	s.engine.Use(gin.Recovery())

	err := s.engine.SetTrustedProxies([]string{"127.0.0.1/32"})
	if err != nil {
		return fmt.Errorf("failed to set trusted proxies: %w", err)
	}

	s.router.Route()

	err = s.engine.Run(":8080")
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
