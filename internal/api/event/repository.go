package events

import (
	"context"
	"errors"
	"events/internal/types"
	"events/internal/utils"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type repositoryMethods interface {
	create(e *Event) error
	getById(id primitive.ObjectID) (e *Event, err *utils.HttpError)
	list() (*[]Event, error)
	update(id primitive.ObjectID, e any) error
}

type repository struct {
	db *mongo.Collection
}

var (
	errorEventUnexistant = errors.New("Event does not exist")
)

func newRepository(db types.Database) *repository {
	return &repository{
		db: db.Collection("events"),
	}
}

func (r repository) create(e *Event) error {
	_, err := r.db.InsertOne(context.TODO(), e)
	return err
}

func (r repository) getById(id primitive.ObjectID) (*Event, *utils.HttpError) {
	filter := bson.D{{Key: "active", Value: true}, {Key: "_id", Value: id}}
	var event *Event

	if err := r.db.FindOne(context.TODO(), filter).Decode(&event); err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		return nil, utils.NewHttpError(errorEventUnexistant, http.StatusNotFound)
	} else if err != nil {
		return nil, utils.NewHttpError(err, http.StatusInternalServerError)
	}

	return event, nil
}

func (r repository) list() (*[]Event, error) {
	filter := bson.D{{Key: "active", Value: true}}
	cur, err := r.db.Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	var events = []Event{}

	if err = cur.All(context.TODO(), &events); err != nil {
		return nil, err
	}

	return &events, nil
}

func (r repository) update(id primitive.ObjectID, data any) error {
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: data}}

	if res, err := r.db.UpdateOne(context.TODO(), filter, update); err == nil && res.ModifiedCount != 0 {
		return nil
	} else {
		return err
	}
}
