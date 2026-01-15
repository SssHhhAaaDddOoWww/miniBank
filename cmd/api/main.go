package main

import (
	"fmt"
	"log"
	"os"

	db "github.com/SssHhhAaaDddOoWww/miniBank/internal/database"
	"github.com/SssHhhAaaDddOoWww/miniBank/internal/database/model"
	"github.com/SssHhhAaaDddOoWww/miniBank/internal/routes"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	db.Connect()
	err = db.DB.AutoMigrate(
		&model.Account{},
		&model.Transaction{},
		&model.Transfer{},
		&model.LedgerEntry{},
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println(" Migration completed!")

	router := gin.Default()
	routes.Routes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	log.Fatal(router.Run(":" + port))
}
