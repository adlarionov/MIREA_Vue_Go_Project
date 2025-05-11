package models

import "gorm.io/gorm"

type DbInstance struct {
	DB *gorm.DB
}

type DotEnv struct {
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
}

var Env DotEnv
