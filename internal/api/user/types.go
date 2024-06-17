package user

import (
	"events/internal/types"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Active    bool               `json:"-"`
	Email     string             `json:"email" binding:"required"`
	FullName  string             `json:"fullname" binding:"required"`
	Id        primitive.ObjectID `bson:"_id" json:"id"`
	Password  string             `json:"-" binding:"required"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

func newUser(creds *types.SignupCredentials) *User {
	return &User{
		Active:    true,
		CreatedAt: time.Now(),
		Email:     creds.Email,
		FullName:  creds.Fullname,
		Id:        primitive.NewObjectID(),
		Password:  creds.Password,
		UpdatedAt: time.Now(),
	}
}

// Updates user updated at attribute with the current time
func (u *User) update() {
	u.UpdatedAt = time.Now()
}
