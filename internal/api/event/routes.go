package events

import (
	"events/internal/api/auth"
	"events/internal/types"
)

var route = "/events"

func RegisterRoutes(api *types.APIServer) {
	router := api.APIRouter.Group(route)
	protectedRouter := api.APIRouter.Group(route).Use(auth.Authenticate)

	repository := newRepository(api.DB)
	service := newService(repository)
	handler := newHandler(service)

	router.GET("", handler.getEvents)
	router.GET("/:id", handler.getEventById)

	protectedRouter.POST("", handler.createEvent)
	protectedRouter.PUT("/:id", handler.updateEvent)
	protectedRouter.DELETE("/:id", handler.deleteEvent)
}
