package events

import (
	"context"
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type repositoryMethods interface {
	create(e *Event) error
	delete(id primitive.ObjectID) error
	getById(id primitive.ObjectID) (e *Event, err error, statusCode int)
	list() (*[]Event, error)
	update(id primitive.ObjectID, e any) error
}

type repository struct {
	db *mongo.Collection
}

func newRepository(db *mongo.Database, collection string) *repository {
	return &repository{
		db: db.Collection(collection),
	}
}

func (r repository) create(e *Event) error {
	res, err := r.db.InsertOne(context.TODO(), e)

	if err != nil {
		return err
	}

	if _, ok := res.InsertedID.(primitive.ObjectID); ok {
		return nil
	}

	return errors.New("Event could not be created (Could not parse id)")
}


func (r repository) getById(id primitive.ObjectID) (*Event, error, int) {
	filter := bson.D{{Key: "active", Value: true}, {Key: "_id", Value: id}}
	var event *Event

	if err := r.db.FindOne(context.TODO(), filter).Decode(&event); err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		return nil, errors.New("Event does not exist"), http.StatusNotFound
	} else if err != nil {
		return nil,err, http.StatusInternalServerError
	}

	return event, nil, http.StatusOK
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

func (r repository) delete(id primitive.ObjectID) error {
	return r.update(id, bson.D{{Key: "active", Value: false}})
}