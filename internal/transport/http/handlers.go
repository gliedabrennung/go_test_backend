package http

import (
	"encoding/json"
	"errors"
	"gobackend/internal/entity"
	"gobackend/internal/repo"
	"gobackend/pkg"
	"io"
	"log"
	"net/http"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var req entity.RawUser
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid JSON or unknown field payload received")
		return
	}

	err = pkg.ValidateUser(req.Username, req.Password)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Failed to close request body: %v", err)
		}
	}(r.Body)

	log.Printf("Received request to create account for user %s", req.Username)
	response, err := repo.CreateAccount(r.Context(), req.Username, req.Password)
	if err != nil {
		log.Printf("Failed to create account for user %s: %v", req.Username, err)
		if errors.Is(err, repo.ErrUserAlreadyExists) {
			writeError(w, http.StatusConflict, "User already exists")
			return
		}
		if errors.Is(err, repo.ErrInvalidInput) {
			writeError(w, http.StatusBadRequest, "Invalid input")
			return
		}
		writeError(w, http.StatusInternalServerError, "Error creating account")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(map[string]interface{}{
		"success": false,
		"error":   message,
	})
	if err != nil {
		return
	}
}
