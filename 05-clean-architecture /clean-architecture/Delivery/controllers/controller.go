package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"taskmanager/auth/Domain"
	"taskmanager/auth/Infrastructure"
	"taskmanager/auth/Usecases"
)

type Controller struct {
	taskUseCase    *Usecases.TaskUseCase
	userUseCase    *Usecases.UserUseCase
	authMiddleware *Infrastructure.AuthMiddleware
}

func NewController(
	taskUseCase *Usecases.TaskUseCase,
	userUseCase *Usecases.UserUseCase,
	authMiddleware *Infrastructure.AuthMiddleware,
) *Controller {
	return &Controller{
		taskUseCase:    taskUseCase,
		userUseCase:    userUseCase,
		authMiddleware: authMiddleware,
	}
}

func (c *Controller) HandleRegister(ctx *gin.Context) {
	var req Domain.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, token, err := c.userUseCase.Register(req)
	if err != nil {
		if err == Domain.ErrUsernameTaken {
			ctx.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	ctx.JSON(http.StatusCreated, Domain.AuthResponse{
		Token: token,
		User:  *user,
	})
}

func (c *Controller) HandleLogin(ctx *gin.Context) {
	var req Domain.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, token, err := c.userUseCase.Login(req)
	if err != nil {
		if err == Domain.ErrInvalidCredentials || err == Domain.ErrNotFound {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to authenticate user"})
		return
	}

	ctx.JSON(http.StatusOK, Domain.AuthResponse{
		Token: token,
		User:  *user,
	})
}

func (c *Controller) HandleGetTasks(ctx *gin.Context) {
	userID, err := c.authMiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not identify user"})
		return
	}

	role, _ := ctx.Get("role")
	isAdmin := role == string(Domain.RoleAdmin)

	tasks, err := c.taskUseCase.GetAllTasks(userID, isAdmin)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get tasks"})
		return
	}

	ctx.JSON(http.StatusOK, Domain.TaskResponse{Tasks: tasks})
}

func (c *Controller) HandleGetTask(ctx *gin.Context) {
	userID, err := c.authMiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not identify user"})
		return
	}

	idStr := ctx.Param("id")

	role, _ := ctx.Get("role")
	isAdmin := role == string(Domain.RoleAdmin)

	task, err := c.taskUseCase.GetTask(idStr, userID, isAdmin)
	if err != nil {
		if err == Domain.ErrInvalidID {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID format"})
			return
		}
		if err == Domain.ErrNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get task"})
		return
	}

	ctx.JSON(http.StatusOK, Domain.TaskResponse{Task: task})
}

func (c *Controller) HandleCreateTask(ctx *gin.Context) {
	userID, err := c.authMiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not identify user"})
		return
	}

	var req Domain.CreateTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := c.taskUseCase.CreateTask(req, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	ctx.JSON(http.StatusCreated, Domain.TaskResponse{Task: task})
}

func (c *Controller) HandleUpdateTask(ctx *gin.Context) {
	userID, err := c.authMiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not identify user"})
		return
	}

	idStr := ctx.Param("id")

	var req Domain.UpdateTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, _ := ctx.Get("role")
	isAdmin := role == string(Domain.RoleAdmin)

	updatedTask, err := c.taskUseCase.UpdateTask(idStr, userID, isAdmin, req)
	if err != nil {
		if err == Domain.ErrInvalidID {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID format"})
			return
		}
		if err == Domain.ErrNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	ctx.JSON(http.StatusOK, Domain.TaskResponse{Task: updatedTask})
}

func (c *Controller) HandleDeleteTask(ctx *gin.Context) {
	userID, err := c.authMiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not identify user"})
		return
	}

	idStr := ctx.Param("id")

	role, _ := ctx.Get("role")
	isAdmin := role == string(Domain.RoleAdmin)

	err = c.taskUseCase.DeleteTask(idStr, userID, isAdmin)
	if err != nil {
		if err == Domain.ErrInvalidID {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID format"})
			return
		}
		if err == Domain.ErrNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}
		if err == Domain.ErrUnauthorized {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized to delete this task"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
