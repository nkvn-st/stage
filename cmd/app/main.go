package main

import (
	"net/http"
	"stage/internal/database"
	"stage/internal/handlers"
	"stage/internal/messageservice"

	"github.com/gorilla/mux"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&messageservice.Message{})

	repo := messageservice.NewMessageRepository(database.DB)
	service := messageservice.NewService(repo)
	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/get", handler.GetMessageHandler).Methods("GET")
	router.HandleFunc("/api/post", handler.PostMessageHandler).Methods("POST")
	router.HandleFunc("/api/patch/{id}", handler.PatchMessageHandler).Methods("PATCH")
	router.HandleFunc("/api/delete/{id}", handler.DeleteMessageHandler).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
