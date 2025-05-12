package main

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"taskmanager/auth/Delivery/controllers"
	"taskmanager/auth/Delivery/routers"
	"taskmanager/auth/Infrastructure"
	"taskmanager/auth/Repositories"
	"taskmanager/auth/Usecases"
)

func main() {
	// Setup context
	ctx := context.Background()

	// Load environment variables or use defaults
	mongoURI := "mongodb://localhost:27017"
	if envURI := os.Getenv("MONGODB_URI"); envURI != "" {
		mongoURI = envURI
	}

	jwtSecret := "default-jwt-should-be-set-in-env-this-is-a-backup"
	if envSecret := os.Getenv("JWT_SECRET"); envSecret != "" {
		jwtSecret = envSecret
	} else {
		log.Println("Warning: Using default JWT secret. Set JWT_SECRET environment variable in production.")
	}

	// Setup MongoDB connection
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Printf("Error disconnecting from MongoDB: %v", err)
		}
	}()

	// Ping the database
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}
	log.Println("Connected to MongoDB!")

	// Initialize collections
	taskCollection := client.Database("taskmanager").Collection("tasks")
	userCollection := client.Database("taskmanager").Collection("users")

	// Initialize repositories
	taskRepo := Repositories.NewTaskRepository(taskCollection, ctx)
	userRepo := Repositories.NewUserRepository(userCollection, ctx)

	// Initialize user repository with unique index for usernames
	if err := userRepo.Initialize(); err != nil {
		log.Fatalf("Failed to initialize user repository: %v", err)
	}

	// Initialize infrastructure services
	jwtService := Infrastructure.NewJWTService(jwtSecret)
	passwordService := Infrastructure.NewPasswordService()
	authMiddleware := Infrastructure.NewAuthMiddleware(jwtService)

	// Initialize use cases
	taskUseCase := Usecases.NewTaskUseCase(taskRepo)
	userUseCase := Usecases.NewUserUseCase(userRepo, passwordService, jwtService)

	// Initialize controllers
	controller := controllers.NewController(taskUseCase, userUseCase, authMiddleware)

	// Initialize and setup router
	router := routers.NewRouter(controller, authMiddleware)
	r := router.Setup()

	// Start the server
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
