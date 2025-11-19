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
	err := config.InitConfig()
	if err != nil {
		return
	}
	err = repo.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := repo.CloseDB()
		if err != nil {
			log.Println(err)
		}
	}()
	serverConfig, err := config.ServerConfig()
	if err != nil {
		log.Fatal(err)
	}
	appRouter := routes.SetupRoutes()
	server := app.StartServer(appRouter, serverConfig)
	defer transport.Shutdown(server)
}
