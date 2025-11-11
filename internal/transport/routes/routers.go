package routes

import (
	"gobackend/internal/transport"
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", transport.Auth)
	return mux
}
