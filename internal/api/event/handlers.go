package events

import (
	"events/internal/api/auth"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	serv service
}

func newHandler(serv service) *handler {
	return &handler{serv}
}

func (h handler) getEvents(ctx *gin.Context) {
	events, err := h.serv.list()

	if err != nil {
    log.Println("Error at getting events:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the events"})
		return
	}

	ctx.JSON(http.StatusOK, events)
}

func (h handler) createEvent(ctx *gin.Context) {
	eventBody := validateRequestBody(ctx)

	if eventBody == (Event{}) {
		return
	}

	eventBody.UserId = auth.GetUserId(ctx)
	event := newEvent(&eventBody)

	if err := h.serv.create(event); err != nil {
		log.Println("Error at creating event:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create the event"})
		return
	}

	ctx.JSON(http.StatusCreated, event)
}

/*
Looks for an event that has the same param id as id.

When called, if the function returns an empty struct, it will handle the errors, when invoking this function
define if whether the event is empty or not (event == (Event{})), if it is empty, only apply a return clause
*/
func (h handler) handleEventExistance(ctx *gin.Context) *Event {
	event, err := h.serv.getById(ctx.Param("id"))

	if err != nil {
		ctx.JSON(err.Status, gin.H{"message": err.Message})
		return &Event{}
	}

	return event
}

func (h handler) getEventById(ctx *gin.Context) {
	if event := h.handleEventExistance(ctx); *event != (Event{}) {
		ctx.JSON(http.StatusOK, event)
	}
}

func (h handler) updateEvent(ctx *gin.Context) {
	event := h.handleEventExistance(ctx)

	if *event == (Event{}) {
		return
	}

	fieldsToUpdate := validateRequestBody(ctx)

	if fieldsToUpdate == (Event{}) {
		return
	}

	event.Name = fieldsToUpdate.Name
	event.Description = fieldsToUpdate.Description
	event.Location = fieldsToUpdate.Location
	event.Date = fieldsToUpdate.Date

	if err := h.serv.update(auth.GetUserId(ctx), event); err != nil {
		ctx.JSON(err.Status, gin.H{"message": err.Message})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func (h handler) deleteEvent(ctx *gin.Context) {
	event := h.handleEventExistance(ctx)

	if *event == (Event{}) {
		return
	}

	if err := h.serv.delete(event.Id, event.UserId, auth.GetUserId(ctx)); err != nil {
		ctx.JSON(err.Status, gin.H{"message": err.Message})
		return
	}
  
	ctx.JSON(http.StatusOK, event)
}
