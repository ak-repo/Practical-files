package service

import (
	"errors"
	"jwt-gin-gorm/internals/common/utils"
	"jwt-gin-gorm/internals/user/models"
	"jwt-gin-gorm/internals/user/repository"
	"jwt-gin-gorm/pkg/jwt_pkg"
	"time"
)

type UserService interface {
	Register(input models.InputUser) (*models.User, error)
	Login(input models.InputUser) (
		*models.Response, error)
	RefreshToken(refresh string) (string, error)
	GetRefreshExpiration() time.Duration
}

type userService struct {
	repo              repository.UserRepository
	jwtSecret         string
	accessExpiration  time.Duration
	refreshExpiration time.Duration
}

func NewUserService(repo repository.UserRepository, jwtSecret string, accessExpiration, refreshExpiration time.Duration) UserService {
	return &userService{repo: repo, jwtSecret: jwtSecret, accessExpiration: accessExpiration, refreshExpiration: refreshExpiration}
}

func (s *userService) Register(input models.InputUser) (*models.User, error) {

	// email and password validation
	if err := input.Validate(); err != nil {
		return nil, err
	}

	// Check email already registered
	exists, err := s.repo.UserExists(input.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("email already registered")
	}

	// Password hash
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	return s.repo.CreateUser(input.Email, hashedPassword)

}

func (s *userService) Login(input models.InputUser) (*models.Response, error) {

	// Check user input valid
	user, err := s.repo.GetUserByEmail(input.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	if !utils.CompareHashAndPassword(input.Password, user.PasswordHash) {
		return nil, errors.New("invalid credentials")
	}

	// Generate JWT token
	accessToken, err := jwt_pkg.GenerateAccessToken(user.ID, user.Role, user.Email, s.jwtSecret, s.accessExpiration)
	if err != nil {
		return nil, err
	}
	refreshToken, err := jwt_pkg.GenerateRefreshToken(user, s.jwtSecret, s.refreshExpiration)
	if err != nil {
		return nil, err
	}

	res := models.Response{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return &res, nil
}

func (s *userService) RefreshToken(refresh string) (string, error) {

	// Validate Refresh token
	claims, err := jwt_pkg.ValidateToken(refresh, s.jwtSecret)
	if err != nil {
		return "", err
	}

	return jwt_pkg.GenerateAccessToken(claims.UserID, claims.Role, claims.Email, s.jwtSecret, s.accessExpiration)
}




func (s *userService) GetRefreshExpiration() time.Duration {
	return s.refreshExpiration
}
