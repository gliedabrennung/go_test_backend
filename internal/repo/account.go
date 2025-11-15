package repo

import (
	"context"
	"errors"
	"fmt"
	"gobackend/internal/entity"
	"gobackend/pkg"

	"github.com/lib/pq"
)

var (
	ErrUserAlreadyExists = fmt.Errorf("user already exists")
	ErrInvalidInput      = fmt.Errorf("invalid input")
)

func CreateAccount(ctx context.Context, username string, password string) (*entity.User, error) {
	gorm := GetDB()
	if err := gorm.WithContext(ctx).Where("username = ?", username).First(&entity.User{}).Error; err == nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, ErrUserAlreadyExists
		}
		if errors.Is(err, ErrUserAlreadyExists) {
			return nil, ErrUserAlreadyExists
		}
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
