package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"taskmanager/auth/data"
	"taskmanager/auth/middleware"
	"taskmanager/auth/models"
)

func HandleGetTasks(c *gin.Context) {
	userID, err := middleware.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not identify user"})
		return
	}

	// Check if user is admin
	role, _ := c.Get("role")
	isAdmin := role == string(models.RoleAdmin)

	// Get tasks based on user role
	tasks, err := data.GetAllTasks(userID, isAdmin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get tasks"})
		return
	}

	c.JSON(http.StatusOK, models.TaskResponse{Tasks: tasks})
}

func HandleGetTask(c *gin.Context) {
	userID, err := middleware.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not identify user"})
		return
	}

	idStr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID format"})
		return
	}

	// Check if user is admin
	role, _ := c.Get("role")
	isAdmin := role == string(models.RoleAdmin)

	// Get task based on user role
	task, err := data.GetTaskByID(id, userID, isAdmin)
	if err != nil {
		if err.Error() == "task not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get task"})
		return
	}

	c.JSON(http.StatusOK, models.TaskResponse{Task: task})
}

func HandleCreateTask(c *gin.Context) {
	userID, err := middleware.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not identify user"})
		return
	}

	var req models.CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now()
	task := models.Task{
		ID:          primitive.NewObjectID(),
		Title:       req.Title,
		Description: req.Description,
		Completed:   req.Completed,
		CreatedAt:   now,
		UpdatedAt:   now,
		UserID:      userID,
	}

	err = data.CreateTask(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, models.TaskResponse{Task: &task})
}

func HandleUpdateTask(c *gin.Context) {
	userID, err := middleware.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not identify user"})
		return
	}

	idStr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID format"})
		return
	}

	var req models.UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user is admin
	role, _ := c.Get("role")
	isAdmin := role == string(models.RoleAdmin)

	// Build update document
	updates := make(map[string]interface{})

	if req.Title != "" {
		updates["title"] = req.Title
	}

	if req.Description != "" {
		updates["description"] = req.Description
	}

	if req.Completed != nil {
		updates["completed"] = *req.Completed
	}

	updatedTask, err := data.UpdateTask(id, userID, isAdmin, updates)
	if err != nil {
		if err.Error() == "task not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, models.TaskResponse{Task: updatedTask})
}

func HandleDeleteTask(c *gin.Context) {
	userID, err := middleware.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not identify user"})
		return
	}

	idStr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID format"})
		return
	}

	// Check if user is admin
	role, _ := c.Get("role")
	isAdmin := role == string(models.RoleAdmin)

	err = data.DeleteTask(id, userID, isAdmin)
	if err != nil {
		if err.Error() == "task not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
