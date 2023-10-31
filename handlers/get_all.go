package handlers

import (
	models "api/models/users"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var lastRequestTime time.Time

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
	datas, err := models.GetAll()
	if err != nil {log.Printf("Erro ao obter %v registros!", err)}
	json.NewEncoder(w).Encode(datas)
	w.Header().Add("Content-Type", "application/json")
}