package handlers

import (
	"api/database"
	models "api/models/users"
	"encoding/json"
	"net/http"
	"time"
)


func GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type not application/json!", http.StatusUnsupportedMediaType)
		return
	}
	mensageContext := map[string]any{
		"status": "429",
		"ok": false,
		"mensage": "Blocked by flood",
		"statusText": http.StatusBadGateway,
	}
	if time.Since(lastRequestTime) < 10*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		mensageContextJson, err := json.Marshal(mensageContext)
		if err != nil {panic(err)}
		w.Write(mensageContextJson)
		return
	}
	lastRequestTime = time.Now()

		var user *database.User
		err := json.NewDecoder(r.Body).Decode(&user)
	 if err != nil {w.WriteHeader(http.StatusBadRequest)}
	
	data, err := models.GetUser(user)
	if err != nil {w.WriteHeader(http.StatusNotFound)}

	json.NewEncoder(w).Encode(&data)
	if err != nil {w.WriteHeader(http.StatusInternalServerError)}
	w.Header().Add("Content-Type", "application/json")
}