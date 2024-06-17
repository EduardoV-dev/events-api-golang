package events

import (
	"events/internal/types"
)

func RegisterRoutes(api *types.APIServer) {
	router := api.APIRouter.Group("/events")

	repository := newRepository(api.DB)
	service := newService(repository)
	handler := newHandler(service)

	router.GET("", handler.getEvents)
	router.GET("/:id", handler.getEventById)

	router.POST("", handler.createEvent)
	router.PUT("/:id", handler.updateEvent)
	router.DELETE("/:id", handler.deleteEvent)
}
