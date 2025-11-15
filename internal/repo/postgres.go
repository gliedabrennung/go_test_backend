package repo

import (
	"context"
	"fmt"
	"log"
	"time"

	"gobackend/internal/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	gormDB *gorm.DB
	_      context.Context
)

func InitDB() error {
	//databaseCfg := config.DBConfig()
	dsn := fmt.Sprintf("host=localhost user=postgres password=1488 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Almaty") //databaseCfg.Host,
	//databaseCfg.User,
	//databaseCfg.Pass,
	//databaseCfg.DBName,
	//databaseCfg.Port,
	//databaseCfg.SSLMode,
	//databaseCfg.TimeZone,

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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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
