package repo

import (
	"context"
	"fmt"
	"gobackend/internal/models"
	"gobackend/pkg"
)

var (
	ErrUserAlreadyExists = fmt.Errorf("user already exists")
	ErrInvalidInput      = fmt.Errorf("invalid input")
)

func CreateAccount(ctx context.Context, username string, password string) (*models.Response, error) {
	gorm := GetDB()
	if err := gorm.WithContext(ctx).Where("username = ?", username).First(&models.User{}).Error; err == nil {
		return nil, ErrUserAlreadyExists
	}

	hashedPassword, err := pkg.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %s", err)
	}

	user := &models.User{
		Username:       username,
		HashedPassword: hashedPassword,
	}

	if err := gorm.WithContext(ctx).Create(user).Error; err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	res := &models.Response{
		ID:       user.ID,
		Username: user.Username,
	}

	return res, nil
}
