package database

import (
	"os"

	"github.com/gsdenys/go-fiber-postgres/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// dsnEnvVar is the name of environment variable for the database dsn
const dsnEnvVar string = "DATABASE_DSN"

// DB represents the global database instance.
var DB *gorm.DB

// InitDB initializes the connection to the database using the DSN provided in the environment.
// It sets up a connection to the PostgreSQL database using GORM and the provided DSN.
// If the connection is successful, it assigns the database instance to the global DB variable.
// If there is an error during the connection attempt, it returns the error.
func InitDB() error {
	// Retrieve the DSN (Data Source Name) from the environment variable
	dsn := os.Getenv(dsnEnvVar)

	// Configure the database connection using GORM and PostgreSQL driver
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&models.User{})

	// Assign the database instance to the global DB variable
	DB = db
	return nil
}
