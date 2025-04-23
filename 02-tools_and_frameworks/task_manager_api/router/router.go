package router

import (
	"net/http"
	"taskmanager/api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	router.GET("/tasks", controllers.HandleGetTasks)
	router.GET("/tasks/:id", controllers.HandleGetTask)
	router.PUT("/tasks/:id", controllers.HandleUpdateTask)
	router.POST("/tasks", controllers.HandleCreateTask)
	router.DELETE("/tasks/:id", controllers.HandleDeleteTask)

	return router
}
