package main

import (
	"bank-account-manager/server"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// Configure logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func loadEnvConfig() (string, error) {
	if err := godotenv.Load(".env"); err != nil {
		return "", fmt.Errorf("failed to load .env file: %w", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("PORT environment variable is not set")
	}

	return port, nil
}

func main() {
	port, err := loadEnvConfig()
	if err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	srv := server.Create()
	if err := srv.Listen(port); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
