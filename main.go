package main

import (
	"net/http"
	"task-manager/config"
	"task-manager/routes"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	godotenv.Load()
	// Initialize the database connection
	config.InitDB()

	r := routes.InitRoutes()

	// Start the server
	http.ListenAndServe(":8080", r)
}
