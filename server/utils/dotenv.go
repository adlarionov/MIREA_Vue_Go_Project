package utils

import (
	"log"
	"os"

	common "server/models"

	"github.com/joho/godotenv"
)

func InitDotEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error while initializing .env file!")
	}

	common.Env.DB_HOST = os.Getenv("DB_HOST")
	common.Env.DB_USER = os.Getenv("DB_USER")
	common.Env.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	common.Env.DB_NAME = os.Getenv("DB_NAME")
	common.Env.DB_PORT = os.Getenv("DB_PORT")
}
