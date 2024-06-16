package storage

import (
	"context"
	"events/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
}

func NewDatabase() *Database {
	return &Database{}
}

func (d Database) StartClient() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Envs.MongoURI))

	if err != nil {
		panic("Could not start mongo instance")
	}

	return client.Database(config.Envs.DatabaseName)
}
