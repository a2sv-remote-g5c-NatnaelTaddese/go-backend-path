package Repositories

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"taskmanager/auth/Domain"
)

type TaskRepository struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewTaskRepository(collection *mongo.Collection, ctx context.Context) *TaskRepository {
	return &TaskRepository{
		collection: collection,
		ctx:        ctx,
	}
}

func (r *TaskRepository) GetByID(id primitive.ObjectID, userID primitive.ObjectID, isAdmin bool) (*Domain.Task, error) {
	var task Domain.Task
	var filter bson.M

	// If user is not admin, only return tasks belonging to the user
	if isAdmin {
		filter = bson.M{"_id": id}
	} else {
		filter = bson.M{"_id": id, "user_id": userID}
	}

	err := r.collection.FindOne(r.ctx, filter).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, Domain.ErrNotFound
		}
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) GetAll(userID primitive.ObjectID, isAdmin bool) ([]Domain.Task, error) {
	var tasks []Domain.Task

	var filter bson.M
	// If user is not admin, only return tasks belonging to the user
	if isAdmin {
		filter = bson.M{}
	} else {
		filter = bson.M{"user_id": userID}
	}

	cursor, err := r.collection.Find(r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.ctx)

	if err = cursor.All(r.ctx, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskRepository) Create(task *Domain.Task) error {
	_, err := r.collection.InsertOne(r.ctx, task)
	return err
}

func (r *TaskRepository) Update(id primitive.ObjectID, userID primitive.ObjectID, isAdmin bool, updates map[string]interface{}) (*Domain.Task, error) {
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

	result, err := r.collection.UpdateOne(r.ctx, filter, update)
	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, Domain.ErrNotFound
	}

	// Get the updated task
	var updatedTask Domain.Task
	err = r.collection.FindOne(r.ctx, bson.M{"_id": id}).Decode(&updatedTask)
	if err != nil {
		return nil, err
	}

	return &updatedTask, nil
}

func (r *TaskRepository) Delete(id primitive.ObjectID, userID primitive.ObjectID, isAdmin bool) error {
	var filter bson.M

	// If user is not admin, only delete tasks belonging to the user
	if isAdmin {
		filter = bson.M{"_id": id}
	} else {
		filter = bson.M{"_id": id, "user_id": userID}
	}

	result, err := r.collection.DeleteOne(r.ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return Domain.ErrNotFound
	}

	return nil
}
