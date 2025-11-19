package pkg

import (
	"errors"
	"fmt"
	"gobackend/internal/entity"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateUser(username string, password string) error {
	rawUser := entity.RawUser{
		Username: username,
		Password: password,
	}
	err := validate.Struct(rawUser)
	if err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, e := range validationErrors {
				switch e.Field() {
				case "Username":
					return fmt.Errorf("username must contain only 4 to 14 characters")
				case "Password":
					return fmt.Errorf("password must contain only 8 to 16 characters")
				}
			}
		}
	}
	return err
}
