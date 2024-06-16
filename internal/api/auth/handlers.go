package auth

import (
	"events/internal/api/user"
	"events/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleSignup(ctx *gin.Context) {
	var u user.User

	if err := ctx.ShouldBindJSON(&u); err != nil {
		utils.Log("Error at binding user:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Wrong Request Body"})
		return
	}

	if err := u.Create(); err != nil {
		utils.Log("Error at creating user:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create the user"})
		return
	}

	ctx.JSON(http.StatusCreated, u)
}
