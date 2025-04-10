package main

import (
	"intro/task_03/library_management/controllers"
	"intro/task_03/library_management/services"
)

func main() {
	
	app := controllers.NewLibraryController(services.NewLibraryService())
	app.Start()
}