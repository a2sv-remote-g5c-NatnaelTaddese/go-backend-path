package router

import (
	"net/http"
	"taskmanager/auth/controllers"
	"taskmanager/auth/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	// Public authentication routes
	router.POST("/register", controllers.HandleRegister)
	router.POST("/login", controllers.HandleLogin)

	api := router.Group("/")
	api.Use(middleware.JWTAuth())
	{
		// Routes available to all authenticated users (regular users & admins)
		api.GET("/tasks", controllers.HandleGetTasks)
		api.GET("/tasks/:id", controllers.HandleGetTask)

		// Routes restricted to admin users only
		admin := api.Group("/")
		admin.Use(middleware.RequireAdmin())
		{
			admin.POST("/tasks", controllers.HandleCreateTask)
			admin.PUT("/tasks/:id", controllers.HandleUpdateTask)
			admin.DELETE("/tasks/:id", controllers.HandleDeleteTask)
		}
	}

	return router
}
