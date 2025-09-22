package controllers

import (
	"jwt-golang-auth-akhil-sharma/database"
	"jwt-golang-auth-akhil-sharma/models"
	"jwt-golang-auth-akhil-sharma/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	var body struct {
		Username string
		Password string
	}

	// Getting out from body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported data info"})
		return
	}

	//validation
	if len(body.Password) < 9 {
		c.JSON(http.StatusBadGateway, gin.H{"error": "password is short"})
		return
	}
	if len(body.Username) < 9 {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Username is short"})
		return
	}

	//check username is already in DB
	if err := database.DB.First(&models.User{}).Where("username=?", body.Username).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already taken"})
		return
	}

	//password hashing
	hash, err := utils.HashPassword(body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password hashing failed"})
		return
	}

	user := models.User{
		Username: body.Username,
		Password: hash,
	}

}

func Login(c *gin.Context) {

}

func GetUser(c *gin.Context) {

	//get id
	param_id := c.Param("user_id")

	// check user have permission to get info
	uid := c.GetString("uid")
	if uid != param_id {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized access"})
		return
	}

	//get user infos
	var user models.User
	if err := database.DB.First(&user).Where("uid=?", uid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found in DB"})
		return

	}

	//response
	c.JSON(http.StatusOK, gin.H{"data": user})

}

func GetUsers(c *gin.Context) {

}
