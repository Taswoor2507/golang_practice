package routes

import (
	handler "golang_project_one/handlers"
	loggerMiddleware "golang_project_one/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Golang Project"))
	}).Methods("GET")
	r.Handle("/items", loggerMiddleware.LoggerMiddleware(http.HandlerFunc(handler.GetAllItems))).Methods("GET")
	r.Handle("/items/{id}", loggerMiddleware.LoggerMiddleware(http.HandlerFunc(handler.GetItemById))).Methods("GET")
	return r
}
