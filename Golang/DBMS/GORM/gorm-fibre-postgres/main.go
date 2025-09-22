package main

import (
	"gorm-fibre-postgres/models"
	"gorm-fibre-postgres/storage"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_books", r.CreateBook)
	api.Delete("/delete_book/:id", r.DeleteBook)
	api.Delete("/get_book/:id", r.GetbookById)
	api.Get("/books", r.GetBooks)
}

func (r *Repository) CreateBook(c *fiber.Ctx) error {

	book := Book{}

	if err := c.BodyParser(&book); err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"error": "request failed"})
		return err
	}
	if err := r.DB.Create(&book).Error; err != nil {
		c.Status(http.StatusBadGateway).JSON(&fiber.Map{"error": "cpoud not create book in DB"})
		return err
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "book added"})

	return nil

}

func (r *Repository) GetBooks(c *fiber.Ctx) error {
	bookModel := []models.Book{}

	if err := r.DB.Find(&bookModel).Error; err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"error": "no books"})
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "book fetched successfully", "data": bookModel})

	return nil

}

func (r *Repository) DeleteBook(c *fiber.Ctx) error {
	bookModel := models.Book{}
	id := c.Params("id")
	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{"error": "id cannot find id"})
		return nil
	}

	if err := r.DB.Delete(bookModel, id).Error; err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"error": "Could not delete the book"})
		return err

	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "book deleted successfully"})

	return nil
}

func (r *Repository) GetbookById(c *fiber.Ctx) error {
	bookModel := models.Book{}

	id := c.Params("id")
	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{"error": "id cannot find id"})
		return nil
	}

	if err := r.DB.Where("id-?", id).Find(&bookModel).Error; err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"error": "Could not get the book"})
		return err
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"book": bookModel})
	return nil
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// 	DB_HOST=localhost
	// DB_PORT=5432
	// DB_USER=ak
	// DB_PASS=4455@mint
	// DB_SSLMODE=disable
	// DB_DBname=usersdb
	//db
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBname:   os.Getenv("DB_NAME"),
	}
	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal(err)
	}

	//migarte book
	if err := models.MigrateBook(db); err != nil {
		log.Fatal("migration failed")
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()

	r.SetupRoutes(app)
	app.Listen(":8080")

}
