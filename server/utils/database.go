package utils

import (
	"fmt"
	"log"

	constants "server/constants"
	"server/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() *gorm.DB {
	dsn := fmt.Sprintf(
		"postgres://%v:%v@%v:5432/%v?sslmode=disable",
		constants.Env.DB_USER,
		constants.Env.DB_PASSWORD,
		constants.Env.DB_HOST,
		constants.Env.DB_NAME)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error while connecting to Database!")
	}

	migrateErr := DB.AutoMigrate(models.Entities...)

	if migrateErr != nil {
		log.Fatal(migrateErr)
	}

	return DB
}
