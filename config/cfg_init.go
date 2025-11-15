package config

import (
	"gobackend/internal/entity"
	"log"

	"github.com/spf13/viper"
)

func InitConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("") // <- put here your config file directory
	return viper.ReadInConfig()
}

func GetConfig() entity.Database {
	err := InitConfig()
	if err != nil {
		log.Println(err)
	}

	return entity.Database{
		Host:     viper.GetString("DB_HOST"),
		Port:     viper.GetString("DB_PORT"),
		User:     viper.GetString("DB_USER"),
		Pass:     viper.GetString("DB_PASSWORD"),
		DBName:   viper.GetString("DB_NAME"),
		TimeZone: viper.GetString("DB_TIMEZONE"),
	}
}
