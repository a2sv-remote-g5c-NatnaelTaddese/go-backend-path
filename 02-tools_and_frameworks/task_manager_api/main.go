package main

import (
	"taskmanager/api/router"
)

func main() {
	router := router.SetupRouter()
	router.Run(":8080")
}
