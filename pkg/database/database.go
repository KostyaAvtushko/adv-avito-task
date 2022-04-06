package database

import (
	"adv-backend-trainee-assignment/pkg/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=avtushko dbname=advAvito port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("err with connecting")
	}

	DB = db

	db.AutoMigrate(entities.Advertisement{})
}
