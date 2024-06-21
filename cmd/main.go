package main

import (
	"events/cmd/api"
	"events/internal/config"
	"events/internal/storage"
	"events/internal/types"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	setGinMode()

	server := gin.Default()
	ping(server)
  
	db := storage.NewDatabase().StartClient()
	apiRouter := server.Group("/api/v1")

	app := &types.APIServer{
		APIRouter: apiRouter,
		DB:        db,
	}

	api.StartAPI(app)
	server.Run(fmt.Sprintf(":%s", config.Envs.AppPort))
}

func ping(server *gin.Engine) {
	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Server Up!"})
	})
}

func setGinMode() {
	if config.Envs.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
}
