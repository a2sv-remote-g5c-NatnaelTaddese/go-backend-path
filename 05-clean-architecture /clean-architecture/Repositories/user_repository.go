package Repositories

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"taskmanager/auth/Domain"
)

type UserRepository struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewUserRepository(collection *mongo.Collection, ctx context.Context) *UserRepository {
	return &UserRepository{
		collection: collection,
		ctx:        ctx,
	}
}

func (r *UserRepository) Initialize() error {
	// Create unique index for username
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := r.collection.Indexes().CreateOne(r.ctx, indexModel)
	return err
}

func (r *UserRepository) Create(user *Domain.User) error {
	_, err := r.collection.InsertOne(r.ctx, user)
	if mongo.IsDuplicateKeyError(err) {
		return Domain.ErrUsernameTaken
	}
	return err
}

func (r *UserRepository) GetByID(id primitive.ObjectID) (*Domain.User, error) {
	var user Domain.User
	err := r.collection.FindOne(r.ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, Domain.ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByUsername(username string) (*Domain.User, error) {
	var user Domain.User
	err := r.collection.FindOne(r.ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, Domain.ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateLastLogin(id primitive.ObjectID) error {
	update := bson.M{
		"$set": bson.M{"last_login_at": time.Now()},
	}

	result, err := r.collection.UpdateOne(r.ctx, bson.M{"_id": id}, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return Domain.ErrNotFound
	}

	return nil
}