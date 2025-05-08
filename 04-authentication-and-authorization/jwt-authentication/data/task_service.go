package data

import (
	"errors"
	"taskmanager/auth/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetTaskByID(id primitive.ObjectID, userID primitive.ObjectID, isAdmin bool) (*models.Task, error) {
	var task models.Task
	var filter bson.M
	
	// If user is not admin, only return tasks belonging to the user
	if isAdmin {
		filter = bson.M{"_id": id}
	} else {
		filter = bson.M{"_id": id, "user_id": userID}
	}

	err := taskCollection.FindOne(ctx, filter).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("task not found")
		}
		return nil, err
	}

	return &task, nil
}

func GetAllTasks(userID primitive.ObjectID, isAdmin bool) ([]models.Task, error) {
	var tasks []models.Task
	
	var filter bson.M
	// If user is not admin, only return tasks belonging to the user
	if isAdmin {
		filter = bson.M{}
	} else {
		filter = bson.M{"user_id": userID}
	}

	cursor, err := taskCollection.Find(ctx, filter)
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
	_, err := taskCollection.InsertOne(ctx, task)
	return err
}

func UpdateTask(id primitive.ObjectID, userID primitive.ObjectID, isAdmin bool, updates map[string]interface{}) (*models.Task, error) {
	var filter bson.M
	
	// If user is not admin, only update tasks belonging to the user
	if isAdmin {
		filter = bson.M{"_id": id}
	} else {
		filter = bson.M{"_id": id, "user_id": userID}
	}

	// Set updatedAt time
	updates["updated_at"] = time.Now()
	
	update := bson.M{
		"$set": updates,
	}

	result, err := taskCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, errors.New("task not found")
	}

	// Get the updated task
	var updatedTask models.Task
	err = taskCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&updatedTask)
	if err != nil {
		return nil, err
	}

	return &updatedTask, nil
}

func DeleteTask(id primitive.ObjectID, userID primitive.ObjectID, isAdmin bool) error {
	var filter bson.M
	
	// If user is not admin, only delete tasks belonging to the user
	if isAdmin {
		filter = bson.M{"_id": id}
	} else {
		filter = bson.M{"_id": id, "user_id": userID}
	}

	result, err := taskCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}

	return nil
}