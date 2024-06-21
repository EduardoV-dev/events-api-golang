package events

import (
	"errors"
	"events/internal/utils"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type service struct {
	repo repositoryMethods
}

func newService(repo repositoryMethods) service {
	return service{
		repo,
	}
}

var (
	errorUserNotAuthor    = errors.New("Unauthorized to access the event")
	errorNotValidIdFormat = errors.New("Provided id does not have a valid mongo ID format")
)

func (s *service) list() (*[]Event, *utils.HttpError) {
	return s.repo.list()
}

func (s *service) getById(idString string) (*Event, *utils.HttpError) {
	id, err := primitive.ObjectIDFromHex(idString)

	if err != nil {
		return nil, utils.NewHttpError(err, http.StatusInternalServerError)
	}

	return s.repo.getById(id)
}

func (s *service) create(e *Event) *utils.HttpError {
	return s.repo.create(e)
}

func checkEventOwning(userId, userCreatorId primitive.ObjectID) *utils.HttpError {
	if owned := userCreatorId == userId; !owned {
		return utils.NewHttpError(errorUserNotAuthor, http.StatusUnauthorized)
	}

	return nil
}

func (s *service) update(userId primitive.ObjectID, upd *Event) *utils.HttpError {
	if err := checkEventOwning(userId, upd.UserId); err != nil {
		return err
	}

	upd.update()
  err := s.repo.update(upd.Id, upd)
  
	return err
}

func (s *service) delete(id, userCreatorId, userId primitive.ObjectID) *utils.HttpError {
	if err := checkEventOwning(userId, userCreatorId); err != nil {
		return err
	}

	activeAttrs := bson.D{{Key: "active", Value: false}}
  err := s.repo.update(id, activeAttrs)
  
	return err
}
