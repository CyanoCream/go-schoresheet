package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-scoresheet/migration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

const (
	DEBUG_MODE = true
)

var db *gorm.DB

func StartDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUsername, dbPassword, dbName)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	migration.AutoMigrate(db)
	//db.Debug().AutoMigrate(&models.User{}, &models.UserRole{}, &models.Role{}, &models.PermissionRole{}, &models.Permission{}, &middleware.Session{})
}

func GetDB() *gorm.DB {
	if DEBUG_MODE {
		return db.Debug()
	}

	return db
}
