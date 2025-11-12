package config

import (
	"github.com/joho/godotenv"
	"gobackend/internal/entity"
	"log"
	"os"
	"strconv"
)

func Init() error {
	_ = godotenv.Load()
	return nil
}

func ServerConfig() (entity.Server, error) {
	serverPort := os.Getenv("ADDR")
	readTimeout := os.Getenv("READ_TIMEOUT")
	writeTimeout := os.Getenv("WRITE_TIMEOUT")

	readTimeoutInt, err := strconv.Atoi(readTimeout)
	if err != nil {
		log.Println(err)
	}
	writeTimeoutInt, err := strconv.Atoi(writeTimeout)
	if err != nil {
		log.Println(err)
	}

	server := entity.Server{
		Addr:         serverPort,
		ReadTimeout:  readTimeoutInt,
		WriteTimeout: writeTimeoutInt,
	}

	return server, err
}

func DBConfig() entity.Database {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	sslmode := os.Getenv("SSL_MODE")
	timeZone := os.Getenv("TIME_ZONE")

	database := entity.Database{
		Host:     host,
		Port:     port,
		User:     user,
		Pass:     password,
		DBName:   dbname,
		SSLMode:  sslmode,
		TimeZone: timeZone,
	}

	return database
}
