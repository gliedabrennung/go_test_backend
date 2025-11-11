package transport

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Shutdown(server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill)
	<-stop
	log.Println("Server is shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Shutdown error: ", err)
	}
	log.Println("Server is shut down.")
}
