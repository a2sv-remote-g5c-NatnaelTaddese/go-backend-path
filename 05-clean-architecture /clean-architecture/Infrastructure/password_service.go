package Infrastructure

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordService struct {
	hashCost int
}

func NewPasswordService() *PasswordService {
	return &PasswordService{
		hashCost: bcrypt.DefaultCost,
	}
}

func (s *PasswordService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), s.hashCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (s *PasswordService) ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
