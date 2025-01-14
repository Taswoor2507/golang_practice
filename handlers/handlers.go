package handler

import (
	"encoding/json"
	"golang_project_one/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var items []models.Item

// get all items
func GetAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// get item by id
func GetItemById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	for _, item := range items {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	http.Error(w, "Item not found", http.StatusNotFound)
}
