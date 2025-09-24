package main

import (
	"jwt-gin-gorm/config"
	"jwt-gin-gorm/internals/common/middleware"
	"jwt-gin-gorm/internals/common/utils"
	"jwt-gin-gorm/internals/user/handler"
	"jwt-gin-gorm/internals/user/models"
	"jwt-gin-gorm/internals/user/repository"
	"jwt-gin-gorm/internals/user/service"
	"jwt-gin-gorm/pkg/database"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {

	// Load configure

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	// Initialize DB
	db, err := database.NewDatabase(cfg.GetDSN())
	if err != nil {
		log.Fatal("Failed to connect DB: ", err)
	}

	// seeder
	SeedAdmin(db.DB)

	// GIN
	r := gin.Default()

	r.Use(cors.Default()) // allow all origins (dev only!)

	// Initiaize user routes
	userRepo := repository.NewUserRepository(db.DB)
	userService := service.NewUserService(userRepo, cfg.JWT.Secret, cfg.JWT.AccessExpiry, cfg.JWT.RefreshExpiry)
	userHandler := handler.NewUserHandler(userService)

	//routes
	userRoute := r.Group("/user")
	{
		userRoute.POST("/register", userHandler.RegistrationHandler)
		userRoute.POST("/login", userHandler.LoginHandler)
		userRoute.POST("/refresh_token", userHandler.RefreshTokenHandler)

		// Protected routes
		userProtected := userRoute.Group("/protected")
		userProtected.Use(middleware.AuthMiddleware(cfg.JWT.Secret))
		{
			userProtected.GET("/", userHandler.ProtectedHandler)

		}
	}

	// Server
	log.Printf("Starting server at %s", cfg.ServerAddress())
	if err := r.Run(cfg.ServerAddress()); err != nil {
		log.Fatal("Server failed to start: ", err)
	}

}

func SeedAdmin(db *gorm.DB) {

	password, _ := utils.HashPassword("1234")
	admin := models.User{
		Email:        "admin@example.com",
		PasswordHash: password,
		Role:         "admin",
	}
	db.FirstOrCreate(&admin, models.User{Email: admin.Email})
}
