package utils

import (
	"fmt"
	"log"

	common "server/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		common.Env.DB_HOST,
		common.Env.DB_USER,
		common.Env.DB_PASSWORD,
		common.Env.DB_NAME,
		common.Env.DB_PORT)

	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatal("Error while connecting to Database!")
	}

	return db
}
