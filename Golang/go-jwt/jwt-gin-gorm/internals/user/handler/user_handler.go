package handler

import (
	"jwt-gin-gorm/internals/user/models"
	"jwt-gin-gorm/internals/user/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {

	return &UserHandler{userService: userService}
}

// Register method
func (h *UserHandler) RegistrationHandler(ctx *gin.Context) {

	var input models.InputUser
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}

	user, err := h.userService.Register(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Response
	ctx.JSON(http.StatusCreated, gin.H{"message": "user registered", "id": user.ID, "role": user.Role})
}

// Login method
func (h *UserHandler) LoginHandler(ctx *gin.Context) {

	var input models.InputUser
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}

	// JWT
	res, err := h.userService.Login(input)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Refresh token store in http secure cookie
	maxAge := int(h.userService.GetRefreshExpiration())
	ctx.SetCookie(
		"RefreshToken",
		res.RefreshToken,
		maxAge,
		"/",
		"",
		false,
		true,
	)

	ctx.JSON(http.StatusOK, gin.H{
		"token":      res.AccessToken,
		"token_type": "Bearer",
		"role":       res.User.Role,
	})
}

// Refresh JWT token
func (h *UserHandler) RefreshTokenHandler(ctx *gin.Context) {

	refreshToken, err := ctx.Cookie("RefreshToken")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"refresh": false, "error": "no refresh token"})
		return
	}

	token, err := h.userService.RefreshToken(refreshToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"refresh": false, "error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"refresh":      true,
		"access_token": token,
		"token_type":   "Bearer",
	})

}

// Protected routes

func (h *UserHandler) ProtectedHandler(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "user auth success",
	})
}
