package utils

import (
	"log"
	"os"

	constants "server/constants"

	"github.com/joho/godotenv"
)

func InitDotEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error while initializing .env file!")
	}

	// DB params
	constants.Env.DB_HOST = os.Getenv("DB_HOST")
	constants.Env.DB_USER = os.Getenv("DB_USER")
	constants.Env.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	constants.Env.DB_NAME = os.Getenv("DB_NAME")
	constants.Env.DB_PORT = os.Getenv("DB_PORT")
	// App params
	constants.Env.APP_PORT = os.Getenv("APP_PORT")
	constants.Env.APP_JWT_SECRET = os.Getenv("APP_JWT_SECRET")
}
