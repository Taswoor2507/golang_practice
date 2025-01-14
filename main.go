package main

import (
	"golang_project_one/routes"
	"log"
	"net/http"
)

func main() {
	r := routes.RegisterRoutes()
	log.Println("Server started on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
