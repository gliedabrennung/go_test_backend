package main

import (
	"gobackend/internal/app"
	"gobackend/internal/config"
	"gobackend/internal/repo"
	"gobackend/internal/transport"
	"gobackend/internal/transport/routes"
	"log"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatalf("FATAL: failed to initialize config: %v", err)
	}
	err = repo.InitDB()
	if err != nil {
		log.Fatalf("FATAL: failed to initialize database: %v", err)
	}
	defer func() {
		err := repo.CloseDB()
		if err != nil {
			log.Printf("ERROR: failed to close database connection: %v", err)
		}
	}()
	serverConfig, err := config.ServerConfig()
	if err != nil {
		log.Fatalf("FATAL: failed to initialize server config: %v", err)
	}
	appRouter := routes.SetupRoutes()
	server := app.StartServer(appRouter, serverConfig)
	defer transport.Shutdown(server)
}
