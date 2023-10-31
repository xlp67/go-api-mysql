package main

import (

	"api/handlers"
	"net/http"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Post("/update", handlers.UpdateUser)
	r.Post("/insert", handlers.Insert)
	r.Get("/getall", handlers.GetAll)
	r.Get("/get", handlers.GetUser)
	r.Delete("/delete", handlers.DeleteUser)
	http.ListenAndServe(":8080", r)
}