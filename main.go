package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	var messages []Message
	DB.Find(&messages)
	json.NewEncoder(w).Encode(messages)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var message Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	DB.Create(&message)
	json.NewEncoder(w).Encode(message)
}

func PatchHandler(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	var message Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	DB.Model(&Message{}).Where("id = ?", id).Update("is_done", message.IsDone)
	DB.Find(&message, id)
	json.NewEncoder(w).Encode(message)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	DB.Unscoped().Delete(&Message{}, id)
	json.NewEncoder(w).Encode(map[string]string{"message": "Message was deleted"})
}

func main() {
	InitDB()

	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/get", GetHandler).Methods("GET")
	router.HandleFunc("/api/post", PostHandler).Methods("POST")
	router.HandleFunc("/api/patch/{id}", PatchHandler).Methods("PATCH")
	router.HandleFunc("/api/delete/{id}", DeleteHandler).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
