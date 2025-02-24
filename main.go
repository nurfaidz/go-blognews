package main

import (
	"github.com/joho/godotenv"
	"go-blognews/database"
	"go-blognews/router"
	"log"
	"os"
)

func main() {
	err := godotenv.Load("config/.env")

	if err != nil {
		log.Fatal("No file .env found, using environment system variables")
	}

	PORT := os.Getenv("API_PORT")

	if PORT == "" {
		PORT = "8080"
	}

	database.StartDB()
	router.StartServer().Run("0.0.0.0:" + PORT)
}
