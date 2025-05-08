package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"taskmanager/auth/models"
)

var (
	JWTSecret       = "default-jwt-should-be-set-in-env-this-is-a-backupq"
	tokenExpiration = 24 * time.Hour

	ErrAuthHeaderMissing = errors.New("authorization header is required")
	ErrInvalidAuthFormat = errors.New("invalid authorization header format")
	ErrInvalidToken      = errors.New("invalid token")
	ErrInsufficientPerms = errors.New("insufficient permissions")
)

// JWTClaims represents the claims in the JWT token
type JWTClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken creates a new JWT token for the given user
func GenerateToken(user *models.User) (string, error) {
	// Create claims with user information
	claims := JWTClaims{
		UserID:   user.ID.Hex(),
		Username: user.Username,
		Role:     string(user.Role),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken validates the JWT token and returns the claims
func ValidateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}

// extracts the token from the Authorization header
func ExtractTokenFromHeader(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", ErrAuthHeaderMissing
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", ErrInvalidAuthFormat
	}

	return parts[1], nil
}

// middleware to verify JWT tokens
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := ExtractTokenFromHeader(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		claims, err := ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Set claims in context for later use
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		c.Next()
	}
}

// RequireAdmin middleware to check if the user is an admin
func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			return
		}

		if role != string(models.RoleAdmin) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
			return
		}

		c.Next()
	}
}

// extracts the user ID from the context
func GetUserIDFromContext(c *gin.Context) (primitive.ObjectID, error) {
	userIDStr, exists := c.Get("userID")
	if !exists {
		return primitive.ObjectID{}, errors.New("user ID not found in context")
	}

	userID, err := primitive.ObjectIDFromHex(userIDStr.(string))
	if err != nil {
		return primitive.ObjectID{}, errors.New("invalid user ID format")
	}

	return userID, nil
}
