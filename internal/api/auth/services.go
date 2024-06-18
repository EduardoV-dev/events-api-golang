package auth

import (
	"errors"
	"events/internal/api/user"
	"events/internal/utils"
)

type service struct {
	repo user.UserRepository
}

func newService(repo user.UserRepository) service {
	return service{
		repo,
	}
}

func (s service) login(creds *loginCredentials) (string, error) {
	user, err := s.repo.GetByEmail(creds.Email)

	if err != nil {
		return "", err
	}

	if ok := utils.ComparePasswords(user.Password, creds.Password); !ok {
    return "", errors.New("Incorrect Password")
  }

	if token, err := generateToken(authClaims{
		Id:       user.Id,
		FullName: user.FullName,
		Email:    user.Email,
	}); err == nil {
		return token, nil
	} else {
		return "", err
	}
}
