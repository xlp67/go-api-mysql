package handlers

import (
	"api/database"
	models "api/models/users"
	"encoding/json"
	"net/http"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type not application/json!", http.StatusUnsupportedMediaType)
		return
	}
	var data *database.User
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Error to convert content", http.StatusBadRequest)
	}
	err = models.UpdateUser(data)
	if err != nil {
		http.Error(w, "Error to update user", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}