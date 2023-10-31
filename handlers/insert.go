package handlers

import (
	"api/database"
	models "api/models/users"
	"encoding/json"
	"net/http"
)

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type not application/json!", http.StatusUnsupportedMediaType)
		return
	}
	var data database.User
err := json.NewDecoder(r.Body).Decode(&data)
if err != nil {
    http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
    return
}
err = models.InsertUser(&data)
	if err != nil {
		http.Error(w, "Erro ao inserir usuário!", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
