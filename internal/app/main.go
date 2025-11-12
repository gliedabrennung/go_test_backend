package app

import (
	"errors"
	"gobackend/internal/entity"
	"log"
	"net/http"
	"time"
)

func StartServer(appRouter http.Handler, serverConfig entity.Server) *http.Server {
	server := &http.Server{
		Addr:         serverConfig.Addr,
		Handler:      appRouter,
		ReadTimeout:  time.Duration(serverConfig.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(serverConfig.WriteTimeout) * time.Second,
	}
	go func() {
		log.Printf("Server is running on %s", server.Addr)
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Listen and Serve error: %v", err)
		}
	}()
	return server
}
