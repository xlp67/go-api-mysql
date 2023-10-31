package handlers

import (
	"api/database"
	models "api/models/users"
	"encoding/json"
	"net/http"
	"time"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// if w.Header().Get("content-type") != "application/json" {
	// 	w.WriteHeader(http.StatusUnsupportedMediaType)
	// }
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
	if err != nil {w.WriteHeader(http.StatusConflict)}
	_, err = models.Delete(user)
	if err != nil {w.WriteHeader(http.StatusBadRequest)}
	// err = json.NewEncoder(w).Encode(data)
	// if err != nil {w.WriteHeader(http.StatusInternalServerError)}
	// w.Header().Add("Content-Type", "application/json")
}