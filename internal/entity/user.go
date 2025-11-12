package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string `gorm:"uniqueIndex;not null"`
	HashedPassword []byte `gorm:"not null"`
}
