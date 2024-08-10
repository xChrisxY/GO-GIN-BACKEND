package configs

import (
	"os"
	"log"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {

	err := godotenv.Load()

	if err != nil {

		log.Fatal("Error loading .env file")

	}

	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") +")/" + os.Getenv("DB_NAME") +"?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {

		log.Fatal("Failed to connect to the database:", err)

	}

	DB = db
	log.Println("Database connection established")

}