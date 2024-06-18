package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	token = strings.Replace(token, "Bearer ", "", 1)

	if claims, ok := validateToken(token); ok {
		ctx.Set("userId", claims.Id.Hex())
		ctx.Next()
	} else {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
	}
}
