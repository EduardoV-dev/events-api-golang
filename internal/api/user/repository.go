package user

import (
	"context"
	"errors"
	"events/internal/types"
	"events/internal/utils"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	create(u *User) *utils.HttpError
	GetByEmail(email string) (u *User, err *utils.HttpError)
}

type repository struct {
	db types.Table
}

var (
	RepositoryTable     = "users"
	errorUserUnexistant = errors.New("User does not exists")
)

func NewUserRepository(db types.Database) repository {
	return repository{
		db: db.Collection(RepositoryTable),
	}
}

func (r repository) create(u *User) *utils.HttpError {
	_, err := r.db.InsertOne(context.TODO(), u)
	return utils.NewHttpError(err, http.StatusInternalServerError)
}

func (r repository) GetByEmail(email string) (*User, *utils.HttpError) {
	user := new(User)
	filter := bson.D{{Key: "active", Value: true}, {Key: "email", Value: email}}

	if err := r.db.FindOne(context.TODO(), filter).Decode(user); err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		return nil, utils.NewHttpError(errorUserUnexistant, http.StatusNotFound)
	} else if err != nil {
		return nil, utils.NewHttpError(err, http.StatusInternalServerError)
	}

	return user, nil
}
