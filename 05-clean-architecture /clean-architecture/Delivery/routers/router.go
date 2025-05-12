package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"taskmanager/auth/Delivery/controllers"
	"taskmanager/auth/Infrastructure"
)

type Router struct {
	controller     *controllers.Controller
	authMiddleware *Infrastructure.AuthMiddleware
}

func NewRouter(controller *controllers.Controller, authMiddleware *Infrastructure.AuthMiddleware) *Router {
	return &Router{
		controller:     controller,
		authMiddleware: authMiddleware,
	}
}

func (r *Router) Setup() *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	// Public authentication routes
	router.POST("/register", r.controller.HandleRegister)
	router.POST("/login", r.controller.HandleLogin)

	// Protected routes
	api := router.Group("/")
	api.Use(r.authMiddleware.JWTAuth())
	{
		// Routes available to all authenticated users (regular users & admins)
		api.GET("/tasks", r.controller.HandleGetTasks)
		api.GET("/tasks/:id", r.controller.HandleGetTask)
		api.POST("/tasks", r.controller.HandleCreateTask)
		api.PUT("/tasks/:id", r.controller.HandleUpdateTask)
		api.DELETE("/tasks/:id", r.controller.HandleDeleteTask)
	}

	return router
}
