package events

import (
	"events/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
  Binds the incoming request body to an event variable, if error is produced, will handle the logs, send a 
  JSON response for the error, and return an empty event. When invoking this function define if whether the
  event is empty or not (event == (Event{})), if it is empty, only apply a return clause
*/
func validateRequestBody(ctx *gin.Context) Event {
  var event Event
  
	if err := ctx.ShouldBindJSON(&event); err != nil {
    utils.Log("Error at binding request body:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Wrong Request Body"})
    return Event{}
	}

  return event
}

