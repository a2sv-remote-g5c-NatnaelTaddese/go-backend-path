# Task Manager API with JWT Authentication

A RESTful API for managing tasks built with Go, Gin framework, and JWT authentication for secure access.

## Project Structure

```
jwt-authentication/
├── main.go                 # Entry point of the application
├── controllers/            # HTTP request handlers
│   ├── controller.go       # Task controllers
│   └── auth_controller.go  # Authentication controllers
├── models/                 # Data models
│   ├── task.go             # Task model
│   └── user.go             # User model
├── data/                   # Data access layer
│   ├── database.go         # Database connection
│   ├── task_service.go     # Task data operations
│   └── user_service.go     # User data operations
├── middleware/             # Middleware components
│   └── auth_middleware.go  # JWT authentication middleware
├── router/                 # API routes definition
│   └── router.go           # Routes configuration
├── docs/                   # Documentation
│   └── api_documentation.md # API documentation
└── go.mod                  # Go module definition
```

## Features

- User authentication with JWT tokens
- Role-based access control (admin and regular users)
- Create, read, update, and delete tasks
- RESTful API design
- MongoDB database integration
- JSON responses
- Error handling

## Authentication System

- **JWT-based authentication**: Secure API access using JSON Web Tokens
- **Role-based access control**: 
  - Admin users: Can perform all operations (GET, POST, PUT, DELETE)
  - Regular users: Can only perform read operations (GET)
- **User registration and login endpoints**
- **Token validation middleware** for protected routes

## API Endpoints

### Authentication Endpoints

| Method | Endpoint     | Description      | Access      |
|--------|-------------|------------------|------------|
| POST   | /register   | Register new user | Public     |
| POST   | /login      | User login       | Public     |

### Task Endpoints

| Method | Endpoint    | Description       | Access                 |
|--------|-------------|-------------------|------------------------|
| GET    | /health     | Health check      | Public                 |
| GET    | /tasks      | List user's tasks | Authenticated          |
| GET    | /tasks/:id  | Get a single task | Authenticated          |
| POST   | /tasks      | Create a task     | Admin only             |
| PUT    | /tasks/:id  | Update a task     | Admin only             |
| DELETE | /tasks/:id  | Delete a task     | Admin only             |

## Getting Started

1. Clone the repository
2. Navigate to the project directory
3. Make sure MongoDB is running locally or set the `MONGODB_URI` environment variable
4. For production, set the `JWT_SECRET` environment variable (defaults to a test value otherwise)
5. Run the application:
   ```
   go run main.go
   ```
6. The API will be available at `http://localhost:8080`

## Authentication Flow

1. Register a user:
   ```bash
   curl -X POST http://localhost:8080/register \
     -H "Content-Type: application/json" \
     -d '{"username": "user1", "password": "password123", "role": "user"}'
   ```

2. Login to get a JWT token:
   ```bash
   curl -X POST http://localhost:8080/login \
     -H "Content-Type: application/json" \
     -d '{"username": "user1", "password": "password123"}'
   ```

3. Use the token in authenticated requests:
   ```bash
   curl -X GET http://localhost:8080/tasks \
     -H "Authorization: Bearer <your-jwt-token>"
   ```

## User Roles

- **Admin**: Create, read, update, and delete any task
- **User**: Read tasks created by themselves

## Data Models

### User Model

| Field        | Type      | Description                    |
|-------------|-----------|--------------------------------|
| id          | ObjectID  | Unique identifier              |
| username    | string    | User's unique username         |
| password    | string    | Hashed password (not returned) |
| role        | string    | User role (admin/user)         |
| created_at  | timestamp | User creation time             |
| last_login_at | timestamp | Last login time              |

### Task Model

| Field       | Type       | Description                  |
|-------------|------------|------------------------------|
| id          | ObjectID   | Unique identifier            |
| title       | string     | Task title                   |
| description | string     | Task description             |
| completed   | boolean    | Completion status            |
| created_at  | timestamp  | Task creation time           |
| updated_at  | timestamp  | Last update time             |
| user_id     | ObjectID   | ID of task creator           |

## Documentation

Detailed API documentation is available in the [API documentation file](docs/api_documentation.md)

## Environment Variables

| Name         | Description                      | Default                    |
|-------------|----------------------------------|----------------------------|
| MONGODB_URI  | MongoDB connection string        | mongodb://localhost:27017  |
| JWT_SECRET   | Secret for signing JWT tokens    | your-secret-key (for dev)  |