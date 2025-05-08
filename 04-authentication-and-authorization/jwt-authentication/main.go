package main

import (
	"log"
	"os"
	"taskmanager/auth/data"
	"taskmanager/auth/middleware"
	"taskmanager/auth/router"
)

func main() {
	// Set JWT Secret from environment variable or use default for development
	if envSecret := os.Getenv("JWT_SECRET"); envSecret != "" {
		middleware.JWTSecret = envSecret
	} else {
		log.Println("Warning: Using default JWT secret. Set JWT_SECRET environment variable in production.")
	}

	// Initialize MongoDB connection
	mongoURI := "mongodb://localhost:27017"
	if envURI := os.Getenv("MONGODB_URI"); envURI != "" {
		mongoURI = envURI
	}

	err := data.InitDB(mongoURI)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer data.CloseDB()

	// Setup and run the router
	r := router.SetupRouter()
	r.Run(":8080")
}