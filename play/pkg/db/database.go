package db

import (
	"machine-task/internals/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB(dsn string) error {

	// db connecting
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// migration
	if err := db.AutoMigrate(&model.UserTable{}); err != nil {
		return err
	}

	DB = db

	return nil

}
