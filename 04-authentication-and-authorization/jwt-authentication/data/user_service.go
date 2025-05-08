package data

import (
	"context"
	"errors"
	"time"

	"taskmanager/auth/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var (
	userCollection *mongo.Collection

	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUsernameExists     = errors.New("username already exists")
)

// InitUserCollection initializes the user collection
func InitUserCollection(client *mongo.Client) {
	userCollection = client.Database("taskmanager").Collection("users")

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := userCollection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		panic(err)
	}
}

func RegisterUser(req models.RegisterRequest) (*models.User, error) {
	// Check if username already exists
	var existingUser models.User
	err := userCollection.FindOne(context.Background(), bson.M{"username": req.Username}).Decode(&existingUser)
	if err == nil {
		return nil, ErrUsernameExists
	} else if err != mongo.ErrNoDocuments {
		return nil, err
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Set role to user if not specified
	role := req.Role
	if role == "" {
		role = models.RoleUser
	}

	// Create the user
	now := time.Now()
	user := models.User{
		ID:          primitive.NewObjectID(),
		Username:    req.Username,
		Password:    string(hashedPassword),
		Role:        role,
		CreatedAt:   now,
		LastLoginAt: now,
	}

	_, err = userCollection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func AuthenticateUser(username, password string) (*models.User, error) {
	var user models.User
	err := userCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	update := bson.M{
		"$set": bson.M{"last_login_at": time.Now()},
	}

	_, err = userCollection.UpdateOne(context.Background(), bson.M{"_id": user.ID}, update)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByID(id primitive.ObjectID) (*models.User, error) {
	var user models.User
	err := userCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}
