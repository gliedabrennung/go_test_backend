package repo

import (
	"context"
	"fmt"
	"gobackend/internal/config"
	"log"
	"time"

	"gobackend/internal/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	gormDB *gorm.DB
	ctx    context.Context
)

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

	sqlDB, err := gormDB.DB()
	if err != nil {
		return fmt.Errorf("error with SQL DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	gormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:   logger.Default.LogMode(logger.Info),
		ConnPool: sqlDB,
	})
	if err != nil {
		return fmt.Errorf("error connection to database: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	status := "up"
	if err := sqlDB.PingContext(ctx); err != nil {
		status = "down"
	}
	log.Println(status)

	type User entity.User
	if err := gormDB.AutoMigrate(&User{}); err != nil {
		return fmt.Errorf("error auto migrate: %w", err)
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
