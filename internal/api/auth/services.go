package auth

import (
	"errors"
	"events/internal/api/user"
	"events/internal/utils"
	"net/http"
)

type service struct {
	repo user.UserRepository
}

func newService(repo user.UserRepository) service {
	return service{
		repo,
	}
}

var (
  errorIncorrectPassword = errors.New("Incorrect Password") 
)

func (s service) login(creds *loginCredentials) (string, *utils.HttpError) {
	user, err := s.repo.GetByEmail(creds.Email)

	if err != nil {
		return "", nil
	}

	if ok := utils.ComparePasswords(user.Password, creds.Password); !ok {
    return "", utils.NewHttpError(errorIncorrectPassword, http.StatusUnauthorized)
  }

	if token, err := generateToken(authClaims{
		Id:       user.Id,
		FullName: user.FullName,
		Email:    user.Email,
	}); err == nil {
		return token, nil
	} else {
		return "", utils.NewHttpError(err, http.StatusInternalServerError)
	}
}
