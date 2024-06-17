package config

import (
	"os"

	"github.com/joho/godotenv"
)

type env struct {
	AppPort         string
	DBAuthMechanism string
	DBHost          string
	DBName          string
	DBPassword      string
	DBPort          string
	DBUser          string
	Env             string
	JwtSecret       string
}

var Envs = initEnvs()

func initEnvs() env {
	err := godotenv.Load()

	if err != nil {
		panic("Could not load envs")
	}

	return env{
		AppPort:         getEnv("APP_PORT", "3000"),
		DBAuthMechanism: getEnv("DB_AUTH_MECHANISM", "SCRAM-SHA-256"),
		DBHost:          getEnv("DB_HOST", "localhost"),
		DBName:          getEnv("DB_NAME", "events"),
		DBPassword:      getEnv("DB_PASSWORD", "secret"),
		DBPort:          getEnv("DB_PORT", "27017"),
		DBUser:          getEnv("DB_USER", "root"),
		Env:             getEnv("ENV", "development"),
    JwtSecret: getEnv("JWT_SECRET", "secret"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
