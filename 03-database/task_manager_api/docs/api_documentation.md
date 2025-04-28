# Task Manager API Documentation

The Task Manager API allows you to manage tasks by creating, retrieving, updating, and deleting them. This API follows RESTful principles and uses JSON for data exchange.

## Base URL

```
http://localhost:8080
```

## API Endpoints

### Health Check

**Endpoint:** `GET /health`

Checks if the API is running.

**Response:**
- Status Code: 200 OK
- Response Body: Plain text "OK"

### List All Tasks

**Endpoint:** `GET /tasks`

Retrieves a list of all tasks.

**Response:**
- Status Code: 200 OK
- Content Type: application/json

```json
{
  "tasks": [
    {
      "id": 1,
      "title": "Task 1",
      "description": "Description for Task 1",
      "completed": false
    },
    {
      "id": 2,
      "title": "Task 2",
      "description": "Description for Task 2",
      "completed": true
    }
  ]
}
```

### Get a Single Task

**Endpoint:** `GET /tasks/:id`

Retrieves a specific task by its ID.

**Parameters:**
- `id` (path parameter): The ID of the task

**Response:**
- Status Code: 200 OK
- Content Type: application/json

```json
{
  "task": {
    "id": 1,
    "title": "Task 1",
    "description": "Description for Task 1",
    "completed": false
  }
}
```

**Error Responses:**
- 400 Bad Request: If the ID is not a valid number
  ```json
  { "error": "Invalid task ID format" }
  ```
- 404 Not Found: If the task does not exist
  ```json
  { "error": "task not found" }
  ```

### Create a Task

**Endpoint:** `POST /tasks`

Creates a new task.

**Request Body:**
```json
{
  "title": "New Task",
  "description": "Description for new task",
  "completed": false
}
```

Note: You don't need to provide the ID as it will be auto-generated.

**Response:**
- Status Code: 201 Created
- Content Type: application/json

```json
{
  "task": {
    "id": 4,
    "title": "New Task",
    "description": "Description for new task",
    "completed": false
  }
}
```

**Error Responses:**
- 400 Bad Request: If the request body is malformed
  ```json
  { "error": "error message" }
  ```
- 500 Internal Server Error: If there's a server error while creating the task
  ```json
  { "error": "error message" }
  ```

### Update a Task

**Endpoint:** `PUT /tasks/:id`

Updates an existing task.

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

**Response:**
- Status Code: 200 OK
- Content Type: application/json

```json
{
  "task": {
    "id": 1,
    "title": "Updated Task",
    "description": "Updated description",
    "completed": true
  }
}
```

**Error Responses:**
- 400 Bad Request: If the ID is not a valid number or request body is malformed
  ```json
  { "error": "Invalid task ID format" }
  ```
- 404 Not Found: If the task does not exist
  ```json
  { "error": "task not found" }
  ```
- 500 Internal Server Error: If there's a server error while updating the task
  ```json
  { "error": "error message" }
  ```

### Delete a Task

**Endpoint:** `DELETE /tasks/:id`

Deletes a task.

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
- 400 Bad Request: If the ID is not a valid number
  ```json
  { "error": "Invalid task ID format" }
  ```
- 404 Not Found: If the task does not exist
  ```json
  { "error": "task not found" }
  ```
- 500 Internal Server Error: If there's a server error while deleting the task
  ```json
  { "error": "error message" }
  ```

## Data Models

### Task

Represents a task in the system.

| Field | Type | Description |
|-------|------|-------------|
| id | integer | Unique identifier for the task |
| title | string | Title of the task |
| description | string | Detailed description of the task |
| completed | boolean | Whether the task has been completed |

## Running the API

The API runs on port 8080 by default. You can start it by running:

```bash
go run main.go
```
