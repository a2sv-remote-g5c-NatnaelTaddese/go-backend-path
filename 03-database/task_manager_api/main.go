package main

import (
	"log"
	"os"
	"taskmanager/api/data"
	"taskmanager/api/router"
)

func main() {
	mongoURI := "mongodb://localhost:27017"
	if envURI := os.Getenv("MONGODB_URI"); envURI != "" {
		mongoURI = envURI
	}

	err := data.InitDB(mongoURI)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer data.CloseDB()

	r := router.SetupRouter()
	r.Run(":8080")
}
