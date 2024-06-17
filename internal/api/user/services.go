package user

import (
	"events/internal/types"
	"events/internal/utils"
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
		utils.Log("Error at hashing password:", err.Error())
		return nil, err
	}

	utils.Log("Hashed password:", hashed, creds.Password)

	creds.Password = hashed
	user := newUser(creds)

	utils.Logf("%+v\n", user)

	return user, s.repo.create(user)
}
