package main

import (
	"bank-account-manager/routes"
	"bank-account-manager/server"
	"os"
	"log"

	"github.com/joho/godotenv"
)

// @Title Bank Account Manager API
// @Version 1.0
// @BasePath /api/v1/
// @Description RESTful API endpoints for Bank Account Management
func main() {
	godotenv.Load(".env")
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	server := server.Create()

	routes.ConfigRoutes(server)

	if err := server.Listen(port); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
