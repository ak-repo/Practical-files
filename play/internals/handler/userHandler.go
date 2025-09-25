package handler

import (
	"machine-task/config"
	"machine-task/internals/model"
	"machine-task/pkg/db"
	jwtpkg "machine-task/pkg/jwt_pkg"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {

	input := model.Input{}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported format"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	user := model.UserTable{
		Email:        input.Email,
		HashPassword: string(hashedPassword),
	}

	result := db.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error", "details": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"userID": user.ID})
}

func Login(c *gin.Context) {

	input := model.Input{}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported format"})
		return
	}

	user := model.UserTable{}
	result := db.DB.Where("email = ?", input.Email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found", "details": result.Error.Error()})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	cfg, err := config.Load()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load config"})
		return
	}

	token, err := jwtpkg.GenerateToken(cfg.SecrectKey, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token generation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Logout(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"msg": "logout"})
}
