package repositories

import (
	validator "github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// Repository handles generic database operations.
type Repository struct {
	db *gorm.DB
}

// NewRepository creates a new instance of Repository.
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// CreateEntity creates a new entity in the database after validating the entity data.
func (r *Repository) CreateEntity(entity interface{}) error {
	// Validate the entity data
	if err := validateEntity(entity); err != nil {
		return err // Return validation error
	}

	// Create the entity in the database
	if err := r.db.Create(entity).Error; err != nil {
		return err // Return database error
	}
	return nil
}

// GetEntityByID retrieves an entity by ID from the database.
func (r *Repository) GetEntityByID(entity interface{}, id uint) error {
	if err := r.db.First(entity, id).Error; err != nil {
		return err
	}
	return nil
}

// validateEntity validates the given entity using the Validate function from the validator package.
func validateEntity(entity interface{}) error {
	validate := validator.New()
	return validate.Struct(entity)
}
