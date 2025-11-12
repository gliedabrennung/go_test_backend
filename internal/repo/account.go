package repo

import (
	"context"
	"fmt"
	"gobackend/internal/entity"
	"gobackend/pkg"
)

func CreateAccount(ctx context.Context, username string, password string) (*entity.User, error) {
	gorm := GetDB()

	err := pkg.ValidateUser(username, password)
	if err != nil {
		return nil, err
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
