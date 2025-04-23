# Task Manager API

A simple RESTful API for managing tasks built with Go and Gin framework.

## Project Structure

```
task_manager/
├── main.go               # Entry point of the application
├── controllers/          # HTTP request handlers
│   └── task_controller.go
├── models/               # Data models
│   └── task.go
├── data/                 # Data access layer
│   └── task_service.go
├── router/               # API routes definition
│   └── router.go
├── docs/                 # Documentation
│   └── api_documentation.md
└── go.mod                # Go module definition
```

## Features

- Create, read, update, and delete tasks
- RESTful API design
- JSON responses
- Error handling

## API Endpoints

| Method | Endpoint    | Description       |
|--------|-------------|-------------------|
| GET    | /health     | Health check      |
| GET    | /tasks      | List all tasks    |
| GET    | /tasks/:id  | Get a single task |
| POST   | /tasks      | Create a task     |
| PUT    | /tasks/:id  | Update a task     |
| DELETE | /tasks/:id  | Delete a task     |

## Getting Started

1. Clone the repository
2. Navigate to the project directory
3. Run the application:
   ```
   go run main.go
   ```
4. The API will be available at `http://localhost:8080`

## Documentation

Detailed API documentation is available in two places:
- Local: See the [API documentation file](docs/api_documentation.md)
- Online: [Postman Documentation](https://documenter.getpostman.com/view/34870519/2sB2ixjtd5)

## Task Model

| Field       | Type    | Description                         |
|-------------|---------|-------------------------------------|
| id          | integer | Unique identifier                   |
| title       | string  | Task title                          |
| description | string  | Task description                    |
| completed   | boolean | Completion status (true/false)      |

## Example Request

```bash
# Create a new task
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title": "New Task", "description": "Task description", "completed": false}'
```
