package api

import (
	"events/internal/api/auth"
	"events/internal/api/event"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type APIServer struct {
	apiRouter *gin.RouterGroup
	db        *mongo.Database
}

func NewAPIServer(apiRouter *gin.RouterGroup, db *mongo.Database) *APIServer {
	return &APIServer{
		apiRouter,
		db,
	}
}

func (a *APIServer) StartAPI() {
	events.RegisterRoutes(a.apiRouter, a.db)
	auth.RegisterRoutes(a.apiRouter)
}
