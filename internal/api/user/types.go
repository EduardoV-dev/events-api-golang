package user

import "events/internal/types"

type User struct {
	*types.Entity 
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"-" validate:"required"`
}
