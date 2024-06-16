package events

import (
	"events/internal/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type handler struct {
	serv service
}

func newHandler(s service) *handler {
	return &handler{s}
}

func (h handler) getEvents(ctx *gin.Context) {
	events, err := h.serv.list()

	if err != nil {
		utils.Log("error at getting events", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the events"})
		return
	}

	ctx.JSON(http.StatusOK, events)
}

func (h handler) createEvent(ctx *gin.Context) {
	event := validateRequestBody(ctx)

	if event == (Event{}) {
		return
	}

	if err := h.serv.create(&event); err != nil {
		utils.Log("Error at creating event:", err)
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
	event, err, statusCode := h.serv.getById(ctx.Param("id"))

	if err != nil {
		utils.Log("Error at searching by id:", err)
		ctx.JSON(statusCode, gin.H{"message": err.Error()})
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
	event.UpdatedAt = time.Now()

	if err := h.serv.update(event.Id, event); err != nil {
		utils.Log("Error at updating event:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event"})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func (h handler) deleteEvent(ctx *gin.Context) {
	event := h.handleEventExistance(ctx)

	if *event == (Event{}) {
		return
	}

	if err := h.serv.delete(event.Id); err != nil {
		utils.Log("Error at deleting event:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event"})
		return
	}

	ctx.JSON(http.StatusOK, event)
}
