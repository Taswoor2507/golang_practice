package routes

import (
	handler "golang_project_one/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Golang Project"))
	}).Methods("GET")
	r.HandleFunc("/items", handler.GetAllItems).Methods("GET")
	r.HandleFunc("/items/{id}", handler.GetItemById).Methods("GET")
	return r
}
