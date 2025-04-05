package auth

import (
	"context"
	"errors"
	"hestia/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authService struct {
    db *gorm.DB
}

type AuthService interface {
	Register(ctx context.Context, email, password string) error
	Authenticate(ctx context.Context, email, password string) (*models.User, error)
}

func NewAuthService(db *gorm.DB) AuthService {
    return &authService{db}
}

func (s *authService) Register(ctx context.Context, email, password string) error {
    hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    user := models.User{
        Email:    email,
        HashedPassword: string(hashed),
    }

    return s.db.WithContext(ctx).Create(&user).Error
}

func (s *authService) Authenticate(ctx context.Context, email, password string) (*models.User, error) {
    var user models.User
    if err := s.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password)); err != nil {
        return nil, errors.New("invalid credentials")
    }

    return &user, nil
}