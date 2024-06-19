package storage

import (
	"context"
	"events/internal/config"
	"events/internal/types"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type database struct {
	uri string
}

var (
	dbName    = config.Envs.DBName
	host      = config.Envs.DBHost
	password  = config.Envs.DBPassword
	port      = config.Envs.DBPort
	user      = config.Envs.DBUser
	extraArgs = config.Envs.DBExtraArgs
)

func NewDatabase() *database {
	isDevelopment := config.Envs.Env == "development"
  
	hostPort := fmt.Sprint(":", port)
	mongoPrefix := "mongodb"
  args := ""

	if !isDevelopment {
		mongoPrefix += "+srv"
		hostPort = ""
    args = extraArgs
	}

	URIString := fmt.Sprintf("%s://%s:%s@%s%s/%s", mongoPrefix, user, password, host, hostPort, args)

	return &database{
		uri: URIString,
	}
}

func (d database) StartClient() types.Database {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOpts := options.Client().ApplyURI(d.uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), clientOpts)

	if err != nil {
		log.Println("Error: ", err.Error())
		panic("Could not start mongo instance")
	}

	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}

	log.Println("Successfully Connected to MongoDB!")
  
	return client.Database(dbName)
}
