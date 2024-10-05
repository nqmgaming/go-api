package main

import (
	"github.com/joho/godotenv"
	"go-api/database"
	"go-api/models"
	"go-api/routes"
	"log"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
	database.Connect()

	if err := database.DB.AutoMigrate(&models.User{}, &models.Book{}); err != nil {
		log.Fatalf("Could not migrate the database: %v", err)
	} else {
		log.Println("Database migrated successfully")
	}
	r := routes.SetupRouter()
	if err := r.SetTrustedProxies([]string{}); err != nil {
		return
	}

	log.Fatal(r.Run(":8080"))
}
