package auth

import (
	"events/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type authClaims struct {
	Id       primitive.ObjectID `json:"id"`
	FullName string             `json:"fullname"`
	Email    string             `json:"email"`
	jwt.RegisteredClaims
}


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
	return token.SignedString([]byte(config.Envs.JwtSecret))
}
