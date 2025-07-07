package main

import (
	"fmt"
	"log"
	"net/http"
	"go-chat-app/internal/profile"
	"go-chat-app/pkg/logger"
	"github.com/gorilla/mux"
)

func main() {
	log := logger.New()

	r := mux.NewRouter()

	r.HandleFunc("/profile-photo/{userId}", profile.GetProfilePhotoHandler(log)).Methods("GET")

	port := ":8080"
	log.Printf("Servidor escuchando en puerto %s...\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
