package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	ISBN   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Get All Books
func createBook(w http.ResponseWriter, r *http.Request) {

}
func getBooks(w http.ResponseWriter, r *http.Request) {

}
func getBook(w http.ResponseWriter, r *http.Request) {

}
func updateBook(w http.ResponseWriter, r *http.Request) {

}
func deleteBook(w http.ResponseWriter, r *http.Request) {

}
func main() {
	// Init router
	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
