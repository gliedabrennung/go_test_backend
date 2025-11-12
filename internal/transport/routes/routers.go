package routes

import (
	http2 "gobackend/internal/transport/http"
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /create", http2.CreateAccount)
	return mux
}
