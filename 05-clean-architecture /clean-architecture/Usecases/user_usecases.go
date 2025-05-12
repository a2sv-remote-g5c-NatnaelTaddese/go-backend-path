package Usecases

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"taskmanager/auth/Domain"
	"taskmanager/auth/Infrastructure"
)

type UserUseCase struct {
	userRepo        Domain.UserRepository
	passwordService *Infrastructure.PasswordService
	jwtService      *Infrastructure.JWTService
}

func NewUserUseCase(
	userRepo Domain.UserRepository,
	passwordService *Infrastructure.PasswordService,
	jwtService *Infrastructure.JWTService,
) *UserUseCase {
	return &UserUseCase{
		userRepo:        userRepo,
		passwordService: passwordService,
		jwtService:      jwtService,
	}
}

func (uc *UserUseCase) Register(req Domain.RegisterRequest) (*Domain.User, string, error) {
	// Hash the password
	hashedPassword, err := uc.passwordService.HashPassword(req.Password)
	if err != nil {
		return nil, "", err
	}

	// Set role to user if not specified
	role := req.Role
	if role == "" {
		role = Domain.RoleUser
	}

	// Create the user
	now := time.Now()
	user := &Domain.User{
		ID:          primitive.NewObjectID(),
		Username:    req.Username,
		Password:    hashedPassword,
		Role:        role,
		CreatedAt:   now,
		LastLoginAt: now,
	}

	// Save the user to the repository
	err = uc.userRepo.Create(user)
	if err != nil {
		return nil, "", err
	}

	// Generate JWT token
	token, err := uc.jwtService.GenerateToken(user.ID.Hex(), user.Username, string(user.Role))
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (uc *UserUseCase) Login(req Domain.LoginRequest) (*Domain.User, string, error) {
	// Find user by username
	user, err := uc.userRepo.GetByUsername(req.Username)
	if err != nil {
		if err == Domain.ErrNotFound {
			return nil, "", Domain.ErrInvalidCredentials
		}
		return nil, "", err
	}

	// Verify password
	err = uc.passwordService.ComparePassword(user.Password, req.Password)
	if err != nil {
		return nil, "", Domain.ErrInvalidCredentials
	}

	// Update last login time
	err = uc.userRepo.UpdateLastLogin(user.ID)
	if err != nil {
		return nil, "", err
	}

	// Generate JWT token
	token, err := uc.jwtService.GenerateToken(user.ID.Hex(), user.Username, string(user.Role))
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (uc *UserUseCase) GetUserByID(id string) (*Domain.User, error) {
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, Domain.ErrInvalidID
	}

	return uc.userRepo.GetByID(userID)
}
