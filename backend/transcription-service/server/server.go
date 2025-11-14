package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"transcription-service/domain/persistance"
	"transcription-service/internal/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	router *router.Router
	repo   persistance.TranscriptionRepository
}

func New(e *gin.Engine, r *router.Router, repo persistance.TranscriptionRepository) *Server {
	return &Server{e, r, repo}
}

func (s *Server) Start() error {
	s.engine.Use(gin.Logger())
	s.engine.Use(gin.Recovery())

	frontendURL, exist := os.LookupEnv("FRONTEND_URL")
	if !exist {
		frontendURL = "http://localhost:8080"
	}

	s.engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{frontendURL},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-User-ID"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	err := s.engine.SetTrustedProxies([]string{"127.0.0.1", "::1"})
	if err != nil {
		return fmt.Errorf("failed to set trusted proxies: %w", err)
	}

	s.router.Route()

	port, exist := os.LookupEnv("APP_PORT")
	if !exist {
		log.Printf("APP_PORT not set, defaulting to 3000")
		port = "3000"
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: s.engine,
	}

	serverErrors := make(chan error, 1)

	// Start the server in a goroutine so we can shut it down gracefully
	go func() {
		log.Printf("Starting server on %s", srv.Addr)
		serverErrors <- srv.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err = <-serverErrors:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("server error: %w", err)
		}

	case sig := <-shutdown:
		log.Printf("Received signal %v, starting graceful shutdown", sig)

		// Give outstanding requests 30 seconds to complete
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err = srv.Shutdown(ctx); err != nil {
			log.Printf("Error during server shutdown: %v", err)

			if closeErr := srv.Close(); closeErr != nil {
				return fmt.Errorf("error forcing server close: %w", closeErr)
			}
			return fmt.Errorf("error during graceful shutdown: %w", err)
		}

		// Close the database connection pool
		if err = s.repo.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}

		log.Println("Server stopped gracefully")
	}

	return nil
}
