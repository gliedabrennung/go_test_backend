package main

import (
	"gobackend/internal/app"
	"gobackend/internal/transport"
	"gobackend/internal/transport/routes"
)

func main() {
	appRouter := routes.SetupRoutes()
	server := app.StartServer(appRouter)
	defer transport.Shutdown(server)
}
