//go:build wireinject
// +build wireinject

package server

import (
	"os"
	"transcription-service/controller/healthcontroller"
	"transcription-service/router"

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
)

func InitializeServer() (*Server, error) {
	wire.Build(ProviderSet)

	return nil, nil
}
