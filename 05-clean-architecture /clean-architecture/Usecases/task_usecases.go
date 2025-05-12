package Usecases

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"taskmanager/auth/Domain"
)

type TaskUseCase struct {
	taskRepo Domain.TaskRepository
}

func NewTaskUseCase(taskRepo Domain.TaskRepository) *TaskUseCase {
	return &TaskUseCase{
		taskRepo: taskRepo,
	}
}

func (uc *TaskUseCase) GetTask(id string, userID primitive.ObjectID, isAdmin bool) (*Domain.Task, error) {
	taskID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, Domain.ErrInvalidID
	}

	task, err := uc.taskRepo.GetByID(taskID, userID, isAdmin)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (uc *TaskUseCase) GetAllTasks(userID primitive.ObjectID, isAdmin bool) ([]Domain.Task, error) {
	return uc.taskRepo.GetAll(userID, isAdmin)
}

func (uc *TaskUseCase) CreateTask(req Domain.CreateTaskRequest, userID primitive.ObjectID) (*Domain.Task, error) {
	now := time.Now()
	task := &Domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       req.Title,
		Description: req.Description,
		Completed:   req.Completed,
		CreatedAt:   now,
		UpdatedAt:   now,
		UserID:      userID,
	}

	err := uc.taskRepo.Create(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (uc *TaskUseCase) UpdateTask(id string, userID primitive.ObjectID, isAdmin bool, req Domain.UpdateTaskRequest) (*Domain.Task, error) {
	taskID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, Domain.ErrInvalidID
	}

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

	if len(updates) == 0 {
		// No updates provided
		task, err := uc.taskRepo.GetByID(taskID, userID, isAdmin)
		return task, err
	}

	return uc.taskRepo.Update(taskID, userID, isAdmin, updates)
}

func (uc *TaskUseCase) DeleteTask(id string, userID primitive.ObjectID, isAdmin bool) error {
	taskID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Domain.ErrInvalidID
	}

	return uc.taskRepo.Delete(taskID, userID, isAdmin)
}
