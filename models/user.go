package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// User represents the structure of a user in the application.
type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null" validate:"required,min=3,max=50"`
	Email    string `json:"email" gorm:"unique;not null" validate:"required,email"`
}

// NewUser creates a new User instance with the provided data.
func NewUser(username, email string) *User {
	return &User{
		Username: username,
		Email:    email,
	}
}

// ValidateUser validates the user struct based on the defined tags.
func ValidateUser(user *User) error {
	validate := validator.New()
	return validate.Struct(user)
}

// UserRepository represents a data access object for user-related operations.
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository.
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// GetUserByID retrieves a user by ID from the database.
func (ur *UserRepository) GetUserByID(userID uint) (*User, error) {
	var user User
	result := ur.db.First(&user, userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// CreateUser saves a new user to the database.
func (ur *UserRepository) CreateUser(user *User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
