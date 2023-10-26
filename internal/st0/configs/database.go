package config

import (
	"fmt"
	"log"

	"github.com/MhmoudGit/file-storage/internal/st0/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() {
	configs := Config()
	var err error
	dbUserName := configs["DB_USERNAME"]
	dbPassword := configs["DB_PASSWORD"]
	dbname := configs["DB_NAME"]

	// Create a connection to the PostgreSQL database.
	dsn := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable", dbUserName, dbPassword, dbname)

	// Open db connection
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	} else {
		log.Println("database connected successfully...")
	}
}

func AutoMigrateDb() {
	// Define a slice of model structs that you want to migrate.
	modelsToMigrate := []interface{}{
		&models.User{},
		&models.Storage{},
		&models.File{},
		// Add more model structs here if needed.
	}
	// // AutoMigrate will create tables if they don't exist based on the model structs.
	err := Db.AutoMigrate(modelsToMigrate...)
	if err != nil {
		log.Fatalf("Error migrating database tables: %v", err)
	}
	log.Println("Tables created/updated successfully...")
}

func Close() {
	// Close db
	dbInstance, _ := Db.DB()
	_ = dbInstance.Close()
	log.Println("database is closed successfully...")
}
