package database

import (
	"jwt-golang-auth-akhil-sharma/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBinit() {
	var err error

	db_url := os.Getenv("DB_URL")

	DB, err = gorm.Open(postgres.Open(db_url), &gorm.Config{})

	if err != nil {
		log.Fatal("Error while connecting into DB: ", err)
	}

	if err = DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Error while migrating DB: ", err)
	}

	log.Println("Successfully connected into DB")

}
