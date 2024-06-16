package events

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRoutes(router *gin.RouterGroup, db *mongo.Database) {
	events := router.Group("/events")

	repository := newRepository(db, "events")
	service := newService(repository)
	handler := newHandler(service)

	events.GET("", handler.getEvents)
	events.GET("/:id", handler.getEventById)

	events.POST("", handler.createEvent)
	events.PUT("/:id", handler.updateEvent)
	events.DELETE("/:id", handler.deleteEvent)
}
