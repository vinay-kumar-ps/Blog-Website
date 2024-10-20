package database

import (
	"log"
	"os"

	"github.com.vinay-kumar-ps/blogbackend/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error  load .env file")
	}
	dsn := os.Getenv("DSN")
	database,err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to database")
	}else{
		log.Println("Connected to port 3000")
	}
	DB =database
	database.AutoMigrate(
		&models.User{},
		&models.Blog{},
	)
}
