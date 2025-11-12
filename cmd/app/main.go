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
	config.Init()
	err := repo.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer repo.CloseDB()

	serverConfig, err := config.ServerConfig()
	if err != nil {
		log.Fatal(err)
	}
	appRouter := routes.SetupRoutes()
	server := app.StartServer(appRouter, serverConfig)
	defer transport.Shutdown(server)
}
