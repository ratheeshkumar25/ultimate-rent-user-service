package db

import (
	"fmt"
	"log"

	"github.com/yourorg/ultimate-rent-car/ultimate-rent-user-service/config"
	"github.com/yourorg/ultimate-rent-car/ultimate-rent-user-service/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB establishes a connection to the database using the provided configuration.
func ConnectDB(config *config.Config) *gorm.DB {
	// Use configuration values to establish a database connection
	host := config.Host
	user := config.User
	password := config.Password
	dbname := config.Database
	port := config.Port
	sslmode := config.Sslmode

	//log the values to verify
	log.Printf("DB Config - Host: %s, User: %s,Password :%s, DBName: %s, Port: %s, SSLMode: %s", host, user, password, dbname, port, sslmode)

	// Implement the logic to connect to the database using GORM
	//Construct the DSN (Data Source Name) for PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)

	// Open the database connection
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate the schema for User, Address, and Document models.
	err = DB.AutoMigrate(
		model.User{},
		model.Address{},
		model.Document{},
	)

	// Check for migration errors
	if err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}

	// Return the database connection instance
	return nil
}
