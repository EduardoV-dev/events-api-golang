package main

import (
	"events/cmd/api"
	"events/internal/config"
	"events/internal/storage"
	"events/internal/types"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	if config.Envs.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	server := gin.Default()
	db := storage.NewDatabase().StartClient()

	apiRouter := server.Group("/api/v1")
	app := &types.APIServer{
		APIRouter: apiRouter,
		DB:        db,
	}

	api.StartAPI(app)
	server.Run(fmt.Sprintf(":%s", config.Envs.AppPort))
}
