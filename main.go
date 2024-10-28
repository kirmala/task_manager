package main

import (
	"log"
	"net/http"
	"task-manager/config"
	"task-manager/routes"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// Initialize the database connection
	config.InitDB()

	r := routes.InitRoutes()

	// Start the server
	http.ListenAndServe(":8080", r)
}
