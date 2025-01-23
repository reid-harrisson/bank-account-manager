package main

import (
	"bank-account-manager/routes"
	"bank-account-manager/server"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadPort() (string, error) {
	if err := godotenv.Load(".env"); err != nil {
		return "", fmt.Errorf("failed to load .env file: %w", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("PORT environment variable is not set")
	}

	return port, nil
}

// @Title Bank Account Manager API
// @Version 1.0
// @BasePath /api/v1/
// @Description RESTful API endpoints for Bank Account Management
func main() {
	port, err := loadPort()
	if err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	server := server.Create()

	routes.ConfigRoutes(server)

	if err := server.Listen(port); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
