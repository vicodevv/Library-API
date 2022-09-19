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

// Init books var as a slice Book struct
var books []Book

// Create a book
func createBook(w http.ResponseWriter, r *http.Request) {

}

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {

}

// Get a Book
func getBook(w http.ResponseWriter, r *http.Request) {

}

// Update a Book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// Delete a Book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}
func main() {
	// Init router
	r := mux.NewRouter()

	// Mock Data - @todo - implement Database
	books = append(books, Book{ID: "1", ISBN: "448743", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})

	// Route handles & endpoints
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
