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
	ErrorUserNotAuthor = errors.New("Unauthorized to access the event")
)

func (s *service) list() (*[]Event, error) {
	return s.repo.list()
}

func (s *service) getById(idString string) (*Event, error, int) {
	id, err := primitive.ObjectIDFromHex(idString)

	if err != nil {
		return nil, errors.New("ID is not a valid mongo ID value"), http.StatusInternalServerError
	}

	return s.repo.getById(id)
}

func (s *service) create(e *Event) error {
	return s.repo.create(e)
}

func checkEventOwning(userId, userCreatorId primitive.ObjectID) error {
	if owned := userCreatorId == userId; !owned {
    utils.Log("owned", owned, userCreatorId, userId)
		return ErrorUserNotAuthor
	}

	return nil
}

func (s *service) update(userId primitive.ObjectID, upd *Event) error {
	if err := checkEventOwning(userId, upd.UserId); err != nil {
		return err
	}

	upd.update()
	return s.repo.update(upd.Id, upd)
}

func (s *service) delete(id, userCreatorId, userId primitive.ObjectID) error {
	if err := checkEventOwning(userId, userCreatorId); err != nil {
		return err
	}

	activeAttrs := bson.D{{Key: "active", Value: false}}
	return s.repo.update(id, activeAttrs)
}
