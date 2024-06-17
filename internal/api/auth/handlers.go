package auth

import (
	"errors"
	"events/internal/api/user"
	"events/internal/types"
	"events/internal/utils"
	"fmt"
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
		utils.Log("Error at binding user:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Wrong Request Body"})
		return
	}

	if user, err := h.userServ.Create(creds); err != nil {
		utils.Log("Error at creating user:", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create the user"})
	} else {
		ctx.JSON(http.StatusCreated, user)
	}
}

func (h handler) login(ctx *gin.Context) {
	creds := new(loginCredentials)

	if err := ctx.ShouldBindJSON(creds); err != nil {
		utils.Log("Error at binding user login:", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Wrong Request Body"})
		return
	}

	token, err := h.authServ.login(creds)

	if err != nil && errors.Is(err, user.UserUnexistantError) {
		utils.Log("Error at login:", err.Error())
		ctx.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("User with %s email does not exist", creds.Email)})
		return
	} else if err != nil {
		utils.Log("Error at login:", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not login user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
