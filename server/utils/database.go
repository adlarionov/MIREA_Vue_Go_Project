package utils

import (
	"fmt"
	"log"

	constants "server/constants"
	"server/models/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		constants.Env.DB_HOST,
		constants.Env.DB_USER,
		constants.Env.DB_PASSWORD,
		constants.Env.DB_NAME,
		constants.Env.DB_PORT)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error while connecting to Database!")
	}

	migrateErr := DB.AutoMigrate(&entity.User{}, &entity.Event{})

	if migrateErr != nil {
		log.Fatal(migrateErr)
	}

	return DB
}
