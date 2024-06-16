package events

import (
	"errors"
	"events/internal/types"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type service struct {
	repo repositoryMethods
}

func newService(repo repositoryMethods) service {
	return service{repo}
}

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
	e.Entity = &types.InitialEntityValues
	return s.repo.create(e)
}

func (s *service) update(id primitive.ObjectID, upd *Event) error {
	if _, err, _ := s.repo.getById(id); err != nil {
		return err
	}
  
	return s.repo.update(id, upd)
}

func (s *service) delete(id primitive.ObjectID) error {
	return s.repo.delete(id)
}
