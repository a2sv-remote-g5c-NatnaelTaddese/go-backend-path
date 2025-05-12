# Task Manager API with Clean Architecture

A RESTful API for managing tasks built with Go, Gin framework, and JWT authentication for secure access, following Clean Architecture principles.

## Project Structure

```
taskmanager/
├── Delivery/             # HTTP layer handling incoming requests
│   ├── main.go           # Entry point of the application
│   ├── controllers/      # HTTP request handlers
│   │   └── controller.go # Task and auth controllers
│   ├── routers/          # API routes definition
│   │   └── router.go     # Routes configuration
├── Domain/               # Enterprise business rules
│   └── domain.go         # Entities, interfaces, and errors
├── Infrastructure/       # External tools and frameworks
│   ├── auth_middleware.go # JWT auth middleware
│   ├── jwt_service.go    # JWT token generation and validation
│   └── password_service.go # Password hashing and comparison
├── Repositories/         # Data access implementations
│   ├── task_repository.go # Task data operations
│   └── user_repository.go # User data operations
├── Usecases/             # Application business rules
│   ├── task_usecases.go  # Task business logic
│   └── user_usecases.go  # User and auth business logic
├── docs/                  # Documentation
│   └── api_documentation.md # API documentation
└── go.mod                 # Go module definition
```

## Clean Architecture Layers

This project follows Clean Architecture, organized into the following layers:

1. **Domain Layer** - The core of the application with entities and business rules:

   - Contains business entities (User, Task)
   - Defines repository interfaces
   - Contains domain errors and constants

2. **Use Case Layer** - Application-specific business rules:

   - Implements the application's use cases (tasks and user management)
   - Orchestrates the flow of data to and from entities
   - Enforces business rules specific to the application

3. **Repository Layer** - Data access implementation:

   - Implements the repository interfaces defined in the Domain layer
   - Handles database operations (MongoDB)
   - Translates between domain entities and database models

4. **Infrastructure Layer** - External tools and frameworks:

   - Password hashing service
   - JWT authentication service
   - Authentication middleware

5. **Delivery Layer** - Framework and delivery mechanisms:
   - HTTP controllers
   - Routing
   - Request/response handling with Gin framework
   - Main application entry point

## Features

- Clean Architecture design for maintainability and testability
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
  - Regular users: Can only access their own tasks
- **User registration and login endpoints**
- **Token validation middleware** for protected routes

## API Endpoints

### Authentication Endpoints

| Method | Endpoint  | Description       | Access |
| ------ | --------- | ----------------- | ------ |
| POST   | /register | Register new user | Public |
| POST   | /login    | User login        | Public |

### Task Endpoints

| Method | Endpoint   | Description       | Access                      |
| ------ | ---------- | ----------------- | --------------------------- |
| GET    | /health    | Health check      | Public                      |
| GET    | /tasks     | List user's tasks | Authenticated               |
| GET    | /tasks/:id | Get a single task | Authenticated               |
| POST   | /tasks     | Create a task     | Authenticated               |
| PUT    | /tasks/:id | Update a task     | Authenticated (Owner/Admin) |
| DELETE | /tasks/:id | Delete a task     | Authenticated (Owner/Admin) |

## Getting Started

1. Clone the repository
2. Navigate to the project directory
3. Make sure MongoDB is running locally or set the `MONGODB_URI` environment variable
4. For production, set the `JWT_SECRET` environment variable (defaults to a test value otherwise)
5. Run the application:
   ```
   go run Delivery/main.go
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

- **Admin**: Can access, create, update, and delete any task in the system
- **User**: Can only access, create, update, and delete their own tasks

## Data Models

### User Model

| Field         | Type      | Description                    |
| ------------- | --------- | ------------------------------ |
| id            | ObjectID  | Unique identifier              |
| username      | string    | User's unique username         |
| password      | string    | Hashed password (not returned) |
| role          | string    | User role (admin/user)         |
| created_at    | timestamp | User creation time             |
| last_login_at | timestamp | Last login time                |

### Task Model

| Field       | Type      | Description        |
| ----------- | --------- | ------------------ |
| id          | ObjectID  | Unique identifier  |
| title       | string    | Task title         |
| description | string    | Task description   |
| completed   | boolean   | Completion status  |
| created_at  | timestamp | Task creation time |
| updated_at  | timestamp | Last update time   |
| user_id     | ObjectID  | ID of task creator |

## Documentation

Detailed API documentation is available in the [API documentation file](docs/api_documentation.md)

## Environment Variables

| Name        | Description                   | Default                                           |
| ----------- | ----------------------------- | ------------------------------------------------- |
| MONGODB_URI | MongoDB connection string     | mongodb://localhost:27017                         |
| JWT_SECRET  | Secret for signing JWT tokens | default-jwt-should-be-set-in-env-this-is-a-backup |
| PORT        | Server port                   | 8080                                              |
