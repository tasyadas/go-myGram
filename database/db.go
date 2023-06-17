package database

import (
	"fmt"
	"go-myGram/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	dbPort := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	config := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable", host, user, dbname, dbPort)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("errror connecting to database :", err)
	}

	fmt.Println("sukses koneksi ke database")
	db.Debug().AutoMigrate(models.User{}, models.Comment{}, models.Photo{}, models.Sosmed{})
}

func GetDB() *gorm.DB {
	return db
}
