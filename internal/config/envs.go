package config

import (
	"os"

	"github.com/joho/godotenv"
)

type env struct {
	DatabaseName string
	Env          string
	MongoURI     string
	Port         string
}

var Envs = initEnvs()

func initEnvs() env {
	err := godotenv.Load()

	if err != nil {
		panic("Could not load envs")
	}

	return env{
		DatabaseName: getEnv("DATABASE_NAME", "events"),
		Env:          getEnv("ENV", "development"),
		MongoURI:     getEnv("MONGO_URI", "mongodb://localhost:27017/"),
		Port:         getEnv("PORT", "3000"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
