package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	InitDB()
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		users := []User{}
		DB.Find(&users)
		c.JSON(200, &users)

	})

	router.Run()
}

//db

var DB *gorm.DB

func InitDB() {
	// Db, err = sql.Open("postgres", "user=ak dbname=users_db password=4455@mint sslmode=disable")

	db, err := gorm.Open(postgres.Open("user=ak dbname=users_db password=4455@mint sslmode=disable"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&User{})
	DB = db
	log.Println("deb connected")

}

// models
type User struct {
	gorm.Model
	Name        string `json:"name"`
	Email       string `json:"email"`
	DateOfBirth string `json:"date_of_birth"`
	Place       string `json:"place"`
	Gender      string `json:"gender"`
}

// id |  name   |           email            | date_of_birth |     place     | gender | created_at | updated_at | deleted_at
