//go:build wireinject
// +build wireinject

package server

import (
	"os"
	"transcription-service/controller/healthcontroller"
	"transcription-service/controller/transcribercontroller"
	"transcription-service/internal/config"
	"transcription-service/internal/router"
	"transcription-service/pkg/minio"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func NewGinEngine() *gin.Engine {
	gin.SetMode(os.Getenv("GIN_MODE"))

	return gin.New()
}

var ProviderSet = wire.NewSet(
	NewGinEngine,
	New,
	router.New,
	healthcontroller.New,
	transcribercontroller.New,
	minio.New,
	config.New,
)

func InitializeServer() (*Server, error) {
	wire.Build(ProviderSet)

	return nil, nil
}
