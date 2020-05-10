package main

import (
  // "fmt"
  // "encoding/json"
  "log"
  "net/http"
  // "math/rand"
  // "strconv"
  "github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
  ID          string     `json:"id"`
  Isbn        string     `json:"isbn"`
  Title       string     `json:"title"`
  Price       float64    `json:"price"`
  Author      *Author    `json:"author"`
  Publisher   *Publisher `json:"publisher"`
}

// Author Struct (Model)
type Author struct {
  Firstname   string     `json:"firstname"`
  Lastname    string     `json:"lastname"`
}


// Publisher Struct (Model)
type Publisher struct {
  Name        string     `json:"name"`
}

// Init books var as a slice Book struct
var books []Book

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {

}

// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {

}

// Create New Book
func createBook(w http.ResponseWriter, r *http.Request) {

}

// Update Book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// Delete Book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}


func main(){
  // Init Router
  r := mux.NewRouter()

  // Mock Data - @todo - implement DB
  books = append(
    books, Book{
      ID: "0", Isbn: "020497", Title: "Lord of the Rings Collection",
      Price: 119.99,
      Author: &Author{Firstname: "JRR", Lastname: "Tolkien"},
      Publisher: &Publisher{Name: "MiddleEarth Binding Co."}
    }
  )
  books = append(
    books, Book{
      ID: "1", Isbn: "110498", Title: "Astrology For You", Price: 22.99,
      Author: &Author{Firstname: "Amanda", Lastname: "Tollefson"},
      Publisher: &Publisher{Name: "Santa Clarita Publishing"}
    }
  )
  books = append(
    books, Book{
      ID: "2", Isbn: "110498", Title: "Lord of the Rings vol 2", Price: 12.99,
      Author: &Author{Firstname: "JRR", Lastname: "Tolkien"},
      Publisher: &Publisher{Name: "MiddleEarth Binding Co."}
    }
  )


  // Route Handlers
  r.HandleFunc("/api/books", getBooks).Methods("GET")
  r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
  r.HandleFunc("/api/books", createBook).Methods("POST")
  r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
  r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":8000", r))
}
