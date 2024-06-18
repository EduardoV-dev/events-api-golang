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
	UserId      primitive.ObjectID `bson:"user_id" json:"user_id"`
}

func newEvent(e *Event) *Event {
	now := time.Now()

	return &Event{
		Active:      true,
		CreatedAt:   now,
		Date:        e.Date,
		Description: e.Description,
		Id:          primitive.NewObjectID(),
		Location:    e.Description,
		Name:        e.Name,
		UpdatedAt:   now,
		UserId:      e.UserId,
	}
}

// Updates the UpdateAt attribute with the current time (time.Now())
func (e *Event) update() {
	e.UpdatedAt = time.Now()
}
