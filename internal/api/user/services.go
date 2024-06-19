package user

import (
	"events/internal/types"
	"events/internal/utils"
	"log"
)

type UserService struct {
	repo UserRepository
}

func NewUserServices(repo UserRepository) UserService {
	return UserService{
		repo,
	}
}

func (s UserService) Create(creds *types.SignupCredentials) (*User, error) {
	hashed, err := utils.HashPassword(creds.Password)

	if err != nil {
		log.Println("Error at hashing password:", err.Error())
		return nil, err
	}

	creds.Password = hashed
	user := newUser(creds)

	return user, s.repo.create(user)
}
