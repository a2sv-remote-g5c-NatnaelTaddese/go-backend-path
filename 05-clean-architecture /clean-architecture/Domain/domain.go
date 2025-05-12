package Domain

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Common errors
var (
	ErrNotFound          = errors.New("resource not found")
	ErrInvalidID         = errors.New("invalid ID format")
	ErrInvalidInput      = errors.New("invalid input")
	ErrUsernameTaken     = errors.New("username already taken")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUnauthorized      = errors.New("unauthorized")
	ErrForbidden         = errors.New("forbidden")
)

// Role represents user role
type Role string

// Available roles
const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

// Task entity represents a task in the system
type Task struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Completed   bool               `json:"completed" bson:"completed"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
	UserID      primitive.ObjectID `json:"user_id" bson:"user_id"`
}

// User entity represents a user in the system
type User struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username    string             `json:"username" bson:"username"`
	Password    string             `json:"-" bson:"password"` // Password is not included in JSON responses
	Role        Role               `json:"role" bson:"role"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	LastLoginAt time.Time          `json:"last_login_at" bson:"last_login_at"`
}

// TaskRepository defines the interface for task data operations
type TaskRepository interface {
	GetByID(id primitive.ObjectID, userID primitive.ObjectID, isAdmin bool) (*Task, error)
	GetAll(userID primitive.ObjectID, isAdmin bool) ([]Task, error)
	Create(task *Task) error
	Update(id primitive.ObjectID, userID primitive.ObjectID, isAdmin bool, updates map[string]interface{}) (*Task, error)
	Delete(id primitive.ObjectID, userID primitive.ObjectID, isAdmin bool) error
}

// UserRepository defines the interface for user data operations
type UserRepository interface {
	Create(user *User) error
	GetByID(id primitive.ObjectID) (*User, error)
	GetByUsername(username string) (*User, error)
	UpdateLastLogin(id primitive.ObjectID) error
}

// TaskRequest and Response DTOs
type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type UpdateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   *bool  `json:"completed"`
}

type TaskResponse struct {
	Task  *Task  `json:"task,omitempty"`
	Tasks []Task `json:"tasks,omitempty"`
}

// Auth Request and Response DTOs
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Role     Role   `json:"role"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}