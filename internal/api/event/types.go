package events

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	Active      bool               `json:"-"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	Date        time.Time          `json:"date" binding:"required"`
	Description string             `json:"description" binding:"required"`
	Id          primitive.ObjectID `bson:"_id" json:"id"`
	Location    string             `json:"location" binding:"required"`
	Name        string             `json:"name" binding:"required"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

// Sets default values to meta data attributes
func (e *Event) initialize() {
	e.Active = true
	e.CreatedAt = time.Now()
	e.Id = primitive.NewObjectID()
	e.UpdatedAt = time.Now()
}

// Updates the UpdateAt attribute with the current time (time.Now())
func (e *Event) update() {
	e.UpdatedAt = time.Now()
}
