package auth

import (
	"events/internal/config"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type authClaims struct {
	Id       primitive.ObjectID `json:"id"`
	FullName string             `json:"fullname"`
	Email    string             `json:"email"`
	jwt.RegisteredClaims
}

var (
	jwtSecret = []byte(config.Envs.JwtSecret)
)

func generateToken(claims authClaims) (string, error) {
	jwtClaims := authClaims{
		Id:       claims.Id,
		FullName: claims.FullName,
		Email:    claims.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	return token.SignedString(jwtSecret)
}

func validateToken(tokenString string) (*authClaims, bool) {
	token, err := jwt.ParseWithClaims(tokenString, &authClaims{}, func(token *jwt.Token) (any, error) {
		return jwtSecret, nil
	})

	if err != nil {
		log.Println("Error at validating token", err.Error())
		return nil, false
	}

	if claims, ok := token.Claims.(*authClaims); ok {
		return claims, true
	}

	return nil, false
}

func GetUserId(ctx *gin.Context) primitive.ObjectID {
	userIdReq := ctx.GetString("userId")
	userId, err := primitive.ObjectIDFromHex(userIdReq)

	if err != nil {
		log.Fatal("Could not parse userId to Object ID", err.Error())
	}

	return userId
}
