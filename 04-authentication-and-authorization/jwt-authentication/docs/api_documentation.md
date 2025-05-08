# Task Manager API Documentation

The Task Manager API allows you to manage tasks by creating, retrieving, updating, and deleting them. This API follows RESTful principles and uses JSON for data exchange.

## Base URL

```
http://localhost:8080
```

## Authentication

This API uses JSON Web Tokens (JWT) for authentication. To access protected endpoints, you need to include the JWT token in the Authorization header of your requests:

```
Authorization: Bearer <your_token>
```

### How to get a token

1. Register a new user using the `/register` endpoint
2. Login with your credentials using the `/login` endpoint
3. Both endpoints will return a JWT token that you can use for authenticated requests

### User Roles

The API supports two user roles:
- **Admin**: Can perform all operations (GET, POST, PUT, DELETE)
- **User**: Can only perform read operations (GET)

## API Endpoints

### Health Check

**Endpoint:** `GET /health`

Checks if the API is running.

**Response:**
- Status Code: 200 OK
- Response Body: Plain text "OK"

### Authentication Endpoints

#### Register

**Endpoint:** `POST /register`

Registers a new user.

**Request Body:**
```json
{
  "username": "newuser",
  "password": "password123",
  "role": "user"
}
```

Note: `role` is optional and defaults to "user" if not specified. Possible values: "admin", "user".

**Response:**
- Status Code: 201 Created
- Content Type: application/json

```json
{
  "token": "your_jwt_token_here",
  "user": {
    "id": "60d21b4667d0d8992e610c85",
    "username": "newuser",
    "role": "user",
    "created_at": "2023-09-01T12:00:00Z",
    "last_login_at": "2023-09-01T12:00:00Z"
  }
}
```

**Error Responses:**
- 400 Bad Request: If the request body is malformed
- 409 Conflict: If the username already exists
- 500 Internal Server Error: If there's a server error

#### Login

**Endpoint:** `POST /login`

Authenticates a user and returns a JWT token.

**Request Body:**
```json
{
  "username": "existinguser",
  "password": "password123"
}
```

**Response:**
- Status Code: 200 OK
- Content Type: application/json

```json
{
  "token": "your_jwt_token_here",
  "user": {
    "id": "60d21b4667d0d8992e610c85",
    "username": "existinguser",
    "role": "user",
    "created_at": "2023-09-01T12:00:00Z",
    "last_login_at": "2023-09-01T12:00:00Z"
  }
}
```

**Error Responses:**
- 400 Bad Request: If the request body is malformed
- 401 Unauthorized: If the credentials are invalid
- 500 Internal Server Error: If there's a server error

### Task Endpoints

All task endpoints require authentication via JWT token.

#### List All Tasks

**Endpoint:** `GET /tasks`

Retrieves a list of tasks. For regular users, returns only their own tasks. For admins, returns all tasks.

**Authentication:** Required

**Response:**
- Status Code: 200 OK
- Content Type: application/json

```json
{
  "tasks": [
    {
      "id": "60d21b4667d0d8992e610c85",
      "title": "Task 1",
      "description": "Description for Task 1",
      "completed": false,
      "created_at": "2023-09-01T12:00:00Z",
      "updated_at": "2023-09-01T12:00:00Z",
      "user_id": "60d21b4667d0d8992e610c85"
    },
    {
      "id": "60d21b4667d0d8992e610c86",
      "title": "Task 2",
      "description": "Description for Task 2",
      "completed": true,
      "created_at": "2023-09-01T12:00:00Z",
      "updated_at": "2023-09-01T12:00:00Z",
      "user_id": "60d21b4667d0d8992e610c85"
    }
  ]
}
```

**Error Responses:**
- 401 Unauthorized: If no JWT token is provided or the token is invalid
- 500 Internal Server Error: If there's a server error

#### Get a Single Task

**Endpoint:** `GET /tasks/:id`

Retrieves a specific task by its ID. Regular users can only access their own tasks.

**Authentication:** Required

**Parameters:**
- `id` (path parameter): The ID of the task

**Response:**
- Status Code: 200 OK
- Content Type: application/json

```json
{
  "task": {
    "id": "60d21b4667d0d8992e610c85",
    "title": "Task 1",
    "description": "Description for Task 1",
    "completed": false,
    "created_at": "2023-09-01T12:00:00Z",
    "updated_at": "2023-09-01T12:00:00Z",
    "user_id": "60d21b4667d0d8992e610c85"
  }
}
```

**Error Responses:**
- 400 Bad Request: If the ID is not a valid format
- 401 Unauthorized: If no JWT token is provided or the token is invalid
- 404 Not Found: If the task does not exist or doesn't belong to the user
- 500 Internal Server Error: If there's a server error

#### Create a Task

**Endpoint:** `POST /tasks`

Creates a new task. Only accessible to users with admin role.

**Authentication:** Required (Admin role)

**Request Body:**
```json
{
  "title": "New Task",
  "description": "Description for new task",
  "completed": false
}
```

**Response:**
- Status Code: 201 Created
- Content Type: application/json

```json
{
  "task": {
    "id": "60d21b4667d0d8992e610c87",
    "title": "New Task",
    "description": "Description for new task",
    "completed": false,
    "created_at": "2023-09-01T12:00:00Z",
    "updated_at": "2023-09-01T12:00:00Z",
    "user_id": "60d21b4667d0d8992e610c85"
  }
}
```

**Error Responses:**
- 400 Bad Request: If the request body is malformed
- 401 Unauthorized: If no JWT token is provided or the token is invalid
- 403 Forbidden: If the user doesn't have admin role
- 500 Internal Server Error: If there's a server error

#### Update a Task

**Endpoint:** `PUT /tasks/:id`

Updates an existing task. Only accessible to users with admin role.

**Authentication:** Required (Admin role)

**Parameters:**
- `id` (path parameter): The ID of the task to update

**Request Body:**
```json
{
  "title": "Updated Task",
  "description": "Updated description",
  "completed": true
}
```

Note: All fields in the request body are optional. Only provided fields will be updated.

**Response:**
- Status Code: 200 OK
- Content Type: application/json

```json
{
  "task": {
    "id": "60d21b4667d0d8992e610c85",
    "title": "Updated Task",
    "description": "Updated description",
    "completed": true,
    "created_at": "2023-09-01T12:00:00Z",
    "updated_at": "2023-09-01T12:10:00Z",
    "user_id": "60d21b4667d0d8992e610c85"
  }
}
```

**Error Responses:**
- 400 Bad Request: If the ID is not a valid format or request body is malformed
- 401 Unauthorized: If no JWT token is provided or the token is invalid
- 403 Forbidden: If the user doesn't have admin role
- 404 Not Found: If the task does not exist
- 500 Internal Server Error: If there's a server error

#### Delete a Task

**Endpoint:** `DELETE /tasks/:id`

Deletes a task. Only accessible to users with admin role.

**Authentication:** Required (Admin role)

**Parameters:**
- `id` (path parameter): The ID of the task to delete

**Response:**
- Status Code: 200 OK
- Content Type: application/json

```json
{
  "message": "Task deleted successfully"
}
```

**Error Responses:**
- 400 Bad Request: If the ID is not a valid format
- 401 Unauthorized: If no JWT token is provided or the token is invalid
- 403 Forbidden: If the user doesn't have admin role
- 404 Not Found: If the task does not exist
- 500 Internal Server Error: If there's a server error

## Data Models

### User

Represents a user in the system.

| Field | Type | Description |
|-------|------|-------------|
| id | string | Unique identifier for the user |
| username | string | Username for authentication |
| password | string | User's password (never returned in responses) |
| role | string | User's role ("admin" or "user") |
| created_at | timestamp | When the user was created |
| last_login_at | timestamp | When the user last logged in |

### Task

Represents a task in the system.

| Field | Type | Description |
|-------|------|-------------|
| id | string | Unique identifier for the task |
| title | string | Title of the task |
| description | string | Detailed description of the task |
| completed | boolean | Whether the task has been completed |
| created_at | timestamp | When the task was created |
| updated_at | timestamp | When the task was last updated |
| user_id | string | ID of the user who created the task |

## Running the API

The API runs on port 8080 by default. You can start it by running:

```bash
go run main.go
```

For production deployments, make sure to set these environment variables:
- `MONGODB_URI`: MongoDB connection string
- `JWT_SECRET`: Secret key used to sign JWT tokens