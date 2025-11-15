package repo

import (
	"context"
	"fmt"
	"gobackend/internal/entity"
	"gobackend/pkg"
)

var (
	ErrUserAlreadyExists = fmt.Errorf("user already exists")
	ErrInvalidInput      = fmt.Errorf("invalid input")
)

func CreateAccount(ctx context.Context, username string, password string) (*entity.User, error) {
	gorm := GetDB()
	if err := gorm.Where("name = ?", username).First(&entity.User{}).Error; err == nil {
		return nil, ErrUserAlreadyExists
	}

	err := pkg.ValidateUser(username, password)
	if err != nil {
		return nil, ErrInvalidInput
	}

	hashedPassword, err := pkg.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %s", err)
	}

	user := &entity.User{
		Username:       username,
		HashedPassword: hashedPassword,
	}

	if err := gorm.WithContext(ctx).Create(user).Error; err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}
