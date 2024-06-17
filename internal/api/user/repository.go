package user

import (
	"context"
	"errors"
	"events/internal/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	create(u *User) error
	GetByEmail(email string) (u *User, err error)
}

type repository struct {
	db types.Table
}

var (
	RepositoryTable     = "users"
	UserUnexistantError = errors.New("User does not exists")
)

func NewUserRepository(db types.Database) repository {
	return repository{
		db: db.Collection(RepositoryTable),
	}
}

func (r repository) create(u *User) error {
	_, err := r.db.InsertOne(context.TODO(), u)
	return err
}

func (r repository) GetByEmail(email string) (*User, error) {
	user := new(User)
	filter := bson.D{{Key: "active", Value: true}, {Key: "email", Value: email}}

	if err := r.db.FindOne(context.TODO(), filter).Decode(user); err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		return nil, UserUnexistantError
	} else if err != nil {
		return nil, err
	}

	return user, nil
}
