package user

import (
	"errors"
	"events/internal/types"
	"events/internal/utils"
	"log"
	"net/http"
)

type UserService struct {
	repo UserRepository
}

var (
	errorPasswordHashing  = errors.New("Could not hash user's password")
	errorEmailAddressUsed = errors.New("The email address is already being used by another user")
)

func NewUserServices(repo UserRepository) UserService {
	return UserService{
		repo,
	}
}

func checkUserExistance(repo UserRepository, email string) bool {
  if user, _ := repo.GetByEmail(email); user != nil {
		return true
	} 

	return false
}

func (s UserService) Create(creds *types.SignupCredentials) (*User, *utils.HttpError) {
	if doUserExists := checkUserExistance(s.repo, creds.Email); doUserExists {
		return nil, utils.NewHttpError(errorEmailAddressUsed, http.StatusConflict)
	}

	hashed, err := utils.HashPassword(creds.Password)

	if err != nil {
		log.Println("Error at hashing password:", err.Error())
		return nil, utils.NewHttpError(errorPasswordHashing, http.StatusInternalServerError)
	}


	creds.Password = hashed
	user := newUser(creds)
  
	return user, s.repo.create(user)
}
