package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	InitialEntityValues = Entity{
		Id:        primitive.NewObjectID(),
		Active:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
)

type Entity struct {
	Id        primitive.ObjectID `bson:"_id" json:"id"`
	Active    bool               `json:"-"`
  CreatedAt time.Time          `bson:"created_at" json:"created_at"`
  UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
