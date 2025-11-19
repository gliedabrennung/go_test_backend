package routes

import (
	api "gobackend/internal/transport/http"
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /create-account/", api.CreateAccount)
	return mux
}
