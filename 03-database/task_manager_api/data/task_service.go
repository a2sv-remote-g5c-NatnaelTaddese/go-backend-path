package data

import (
	"context"
	"errors"
	"taskmanager/api/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetTask(id int) (*models.Task, error) {
	var task models.Task
	filter := bson.M{"id": id}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, filter).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("task not found")
		}
		return nil, err
	}

	return &task, nil
}

func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func CreateTask(task *models.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var maxIDTask models.Task
	opts := options.FindOne().SetSort(bson.D{{Key: "id", Value: -1}})
	err := collection.FindOne(ctx, bson.M{}, opts).Decode(&maxIDTask)

	if err != nil && err != mongo.ErrNoDocuments {
		return err
	} else if err == mongo.ErrNoDocuments {
		task.ID = 1
	} else {
		task.ID = maxIDTask.ID + 1
	}

	_, err = collection.InsertOne(ctx, task)
	return err
}

func UpdateTask(id int, title string, description string, completed bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"id": id}
	update := bson.M{
		"$set": bson.M{
			"title":       title,
			"description": description,
			"completed":   completed,
		},
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("task not found")
	}

	return nil
}

func DeleteTask(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"id": id}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}

	return nil
}
