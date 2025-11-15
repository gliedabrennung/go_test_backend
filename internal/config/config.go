package config

import (
	"gobackend/internal/entity"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func Init() error {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}
	return nil
}

func ServerConfig() (entity.Server, error) {
	serverPort := os.Getenv("ADDR")
	if serverPort == "" {
		serverPort = ":8080"
	}

	readTimeout := os.Getenv("READ_TIMEOUT")
	writeTimeout := os.Getenv("WRITE_TIMEOUT")

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

	server := entity.Server{
		Addr:         serverPort,
		ReadTimeout:  time.Duration(readTimeoutInt),
		WriteTimeout: time.Duration(writeTimeoutInt),
	}

	return server, nil
}
