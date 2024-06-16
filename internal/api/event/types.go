package events

import (
	"events/internal/types"
	"time"
)

type Event struct {
	*types.Entity `bson:",inline"`
	Name          string    `json:"name" validate:"required"`
	Description   string    `json:"description" validate:"required"`
	Location      string    `json:"location" validate:"required"`
	Date          time.Time `json:"date" validate:"required"`
}
