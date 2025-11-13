//go:build wireinject
// +build wireinject

package server

import (
	"os"
	"transcription-service/controller/healthcontroller"
	"transcription-service/controller/transcribercontroller"
	"transcription-service/domain/persistance"
	"transcription-service/domain/transcriber"
	"transcription-service/internal/config"
	"transcription-service/internal/router"
	"transcription-service/pkg/faster-whisper"
	"transcription-service/pkg/minio"
	"transcription-service/usecase/transcription"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func NewGinEngine() *gin.Engine {
	gin.SetMode(os.Getenv("GIN_MODE"))

	return gin.New()
}

var ProviderSet = wire.NewSet(
	wire.Bind(new(persistance.Storage), new(*minio.Storage)),
	wire.Bind(new(transcriber.Transcriber), new(*faster_whisper.Client)),
	NewGinEngine,
	New,
	transcription.New,
	healthcontroller.New,
	transcribercontroller.New,
	router.New,
	minio.New,
	faster_whisper.New,
	config.New)

func InitializeServer() (*Server, error) {
	wire.Build(ProviderSet)

	return nil, nil
}
