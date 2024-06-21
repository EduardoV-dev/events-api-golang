package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type env struct {
	AppPort     string
	DBPrefix    string
	DBExtraArgs string
	DBHost      string
	DBName      string
	DBPassword  string
	DBPort      string
	DBUser      string
	Env         string
	JwtSecret   string
}

var Envs = initEnvs()

// Enables passing an env variable when executing go run command in the terminal
// to run the go application, this way, multiple environments could be executed locally
func getEnvFileToLoad() string {
	environment := flag.String("env", "", "environment to load the .env file")
	flag.Parse()

	envFileToLoad := ".env"

	if *environment != "" {
		envFileToLoad += fmt.Sprintf(".%s", *environment)
	}

	return envFileToLoad
}

func initEnvs() env {
	err := godotenv.Load(getEnvFileToLoad())

	if err != nil {
		panic(fmt.Sprintf("Could not load envs: %s", err.Error()))
	}

	return env{
		AppPort:     getEnv("APP_PORT", "3000"),
		DBPrefix:    getEnv("DB_PREFIX", "mongodb"),
		DBExtraArgs: getEnv("DB_EXTRA_ARGS", ""),
		DBHost:      getEnv("DB_HOST", "localhost"),
		DBName:      getEnv("DB_NAME", "events"),
		DBPassword:  getEnv("DB_PASSWORD", "secret"),
		DBPort:      getEnv("DB_PORT", "27017"),
		DBUser:      getEnv("DB_USER", "root"),
		Env:         getEnv("ENV", "development"),
		JwtSecret:   getEnv("JWT_SECRET", "secret"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
