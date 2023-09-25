package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var items []Item

func main() {
	router := mux.NewRouter()
	items = append(items, Item{ID: 1, Name: "Item 1"})
	items = append(items, Item{ID: 2, Name: "Item 2"})

	router.HandleFunc("/", GetItems).Methods("GET")
	router.HandleFunc("/items", GetItems).Methods("GET")
	router.HandleFunc("/items/{id}", GetItem).Methods("GET")
	router.HandleFunc("/items", CreateItem).Methods("POST")
	router.HandleFunc("/items/{id}", UpdateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", DeleteItem).Methods("DELETE")

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, item := range items {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Item{})
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.ID = len(items) + 1
	items = append(items, item)
	json.NewEncoder(w).Encode(item)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, item := range items {
		if item.ID == id {
			var updatedItem Item
			_ = json.NewDecoder(r.Body).Decode(&updatedItem)
			updatedItem.ID = id
			items[index] = updatedItem
			json.NewEncoder(w).Encode(updatedItem)
			return
		}
	}

	json.NewEncoder(w).Encode(&Item{})
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, item := range items {
		if item.ID == id {
			items = append(items[:index], items[index+1:]...)
			break
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
