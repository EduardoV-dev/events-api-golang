package auth

import (
	"events/internal/api/user"
	"events/internal/types"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	userServ user.UserService
	authServ service
}

func newHandler(userServ user.UserService, authServ service) handler {
	return handler{
		userServ,
		authServ,
	}
}

func (h handler) signup(ctx *gin.Context) {
	creds := new(types.SignupCredentials)

	if err := ctx.ShouldBindJSON(creds); err != nil {
		log.Println("Error at binding user:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Wrong Request Body"})
		return
	}

	if user, err := h.userServ.Create(creds); err != nil {
		ctx.JSON(err.Status, gin.H{"message": err.Message})
	} else {
		ctx.JSON(http.StatusCreated, user)
	}
}

func (h handler) login(ctx *gin.Context) {
	creds := new(loginCredentials)

	if err := ctx.ShouldBindJSON(creds); err != nil {
		log.Println("Error at binding user login:", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Wrong Request Body"})
		return
	}

	token, err := h.authServ.login(creds)

	if err != nil && err.Status == http.StatusNotFound {
		ctx.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("User with %s email does not exist", creds.Email)})
		return
	} else if err != nil {
		log.Println("Error at login:", err.Message)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not login user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
