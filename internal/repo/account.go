package repo

import (
	"fmt"
	"gobackend/internal/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}

var gormDB *gorm.DB

func InitDB() error {
	databaseCfg := config.DBConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		databaseCfg.Host,
		databaseCfg.User,
		databaseCfg.Pass,
		databaseCfg.DBName,
		databaseCfg.Port,
		databaseCfg.SSLMode,
		databaseCfg.TimeZone,
	)
	var err error
	gormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return fmt.Errorf("error connection to database: %w", err)
	}
	sqlDB, err := gormDB.DB()
	if err != nil {
		return fmt.Errorf("error with SQL DB: %w", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("error ping database: %w", err)
	}

	return nil
}

func GetDB() *gorm.DB {
	return gormDB
}

func CloseDB() error {
	if gormDB != nil {
		sqlDB, err := gormDB.DB()
		if err != nil {
			return fmt.Errorf("cannot get SQL DB: %w", err)
		}
		return sqlDB.Close()
	}
	return nil
}
