package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	dsn := os.Getenv("DB")
	if dsn == " " {
		fmt.Println("connection string required !!")
	}
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to db !!", err)
	}

	fmt.Println("Connected succesfully !!!")
}
