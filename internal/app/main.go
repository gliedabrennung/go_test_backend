package app

import (
	"errors"
	"log"
	"net/http"
	"time"
)

func StartServer(appRouter http.Handler) *http.Server {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      appRouter,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
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
