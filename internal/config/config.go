package config

import (
	"gobackend/internal/models"
	"log"
	"strconv"

	"github.com/spf13/viper"
)

func InitConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("C:\\Users\\gnida\\GolandProjects\\gobackend\\") // <- put here your config file directory
	return viper.ReadInConfig()
}

func ServerConfig() (models.Server, error) {
	serverPort := viper.GetString("ADDR")
	if serverPort == "" {
		serverPort = ":8080"
	}

	readTimeout := viper.GetString("READ_TIMEOUT")
	writeTimeout := viper.GetString("WRITE_TIMEOUT")

	readTimeoutInt, err := strconv.Atoi(readTimeout)
	if err != nil {
		log.Printf("Error parsing READ_TIMEOUT: %v, using default 5", err)
		readTimeoutInt = 5
	}

	writeTimeoutInt, err := strconv.Atoi(writeTimeout)
	if err != nil {
		log.Printf("Error parsing WRITE_TIMEOUT: %v, using default 5", err)
		writeTimeoutInt = 5
	}

	server := models.Server{
		Addr:         serverPort,
		ReadTimeout:  readTimeoutInt,
		WriteTimeout: writeTimeoutInt,
	}

	return server, nil
}

func GetDatabaseConfig() models.Database {
	return models.Database{
		Host:     viper.GetString("DB_HOST"),
		Port:     viper.GetString("DB_PORT"),
		User:     viper.GetString("DB_USER"),
		Pass:     viper.GetString("DB_PASSWORD"),
		DBName:   viper.GetString("DB_NAME"),
		TimeZone: viper.GetString("DB_TIMEZONE"),
	}
}
