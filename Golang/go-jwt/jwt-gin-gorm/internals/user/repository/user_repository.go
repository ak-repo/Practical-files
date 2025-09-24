package repository

import (
	"errors"
	"jwt-gin-gorm/internals/user/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	UserExists(email string) (bool, error)
	CreateUser(email, passwordHash string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Check if user exists by email
func (r *userRepository) UserExists(email string) (bool, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil // user doesn’t exist
	}
	if err != nil {
		return false, err // DB error
	}
	return true, nil // user exists
}

// Create new user and return ID
func (r *userRepository) CreateUser(email, passwordHash string) (*models.User, error) {
	user := models.User{
		Email:        email,
		PasswordHash: passwordHash,
		Role:         "customer",
	}
	result := r.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// Get user by email
func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
