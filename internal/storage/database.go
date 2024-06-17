package storage

import (
	"context"
	"events/internal/config"
	"events/internal/types"
	"events/internal/utils"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type database struct {
	uri           string
	authMechanism string
	password      string
	user          string
}

var (
	authMechanism = config.Envs.DBAuthMechanism
	dbName        = config.Envs.DBName
	host          = config.Envs.DBHost
	password      = config.Envs.DBPassword
	port          = config.Envs.DBPort
	uri           = fmt.Sprintf("mongodb://%s:%s/", host, port)
	user          = config.Envs.DBUser
)

func NewDatabase() *database {
	return &database{
		uri,
		authMechanism,
		password,
		user,
	}
}

func (d database) StartClient() types.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	credentials := options.Credential{
		AuthMechanism: d.authMechanism,
		Username:      d.user,
		Password:      d.password,
	}

	clientOpts := options.Client().ApplyURI(d.uri).SetAuth(credentials)
	client, err := mongo.Connect(ctx, clientOpts)

	if err != nil {
		utils.Log("Error: ", err)
		panic("Could not start mongo instance")
	}

	return client.Database(dbName)
}
