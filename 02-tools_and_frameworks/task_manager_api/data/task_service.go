package data

import (
	"errors"
	"slices"
	"taskmanager/api/models"
)

var Tasks = []models.Task{
	models.Task{
		ID:          1,
		Title:       "Task 1",
		Description: "Description for Task 1",
		Completed:   false,
	},
	models.Task{
		ID:          2,
		Title:       "Task 2",
		Description: "Description for Task 2",
		Completed:   true,
	},
	models.Task{
		ID:          3,
		Title:       "Task 3",
		Description: "Description for Task 3",
		Completed:   false,
	},
}

func GetTask(id int) (*models.Task, error) {
	for _, task := range Tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, errors.New("task not found")
}

func CreateTask(task *models.Task) error {
	id := len(Tasks) + 1
	task.ID = id
	Tasks = append(Tasks, *task)
	return nil
}

func UpdateTask(id int, title string, description string, completed bool) error {
	for i, task := range Tasks {
		if task.ID == id {
			task.Update(title, description, completed)
			Tasks[i] = task
			return nil
		}
	}
	return errors.New("task not found")
}

func DeleteTask(id int) error {
	for i, task := range Tasks {
		if task.ID == id {
			Tasks = slices.Delete(Tasks, i, i+1)
			return nil
		}
	}
	return errors.New("task not found")
}
