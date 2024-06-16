package main

import (
	"events/cmd/api"
	"events/internal/config"
	"events/internal/storage"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	db := storage.NewDatabase().StartClient()
	apiRouter := server.Group("/api/v1")
	app := api.NewAPIServer(apiRouter, db)
  
	app.StartAPI()
	server.Run(fmt.Sprintf(":%s", config.Envs.Port))
}
