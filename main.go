package main

import (
  // "fmt"
  "encoding/json"
  "log"
  "net/http"
  "math/rand"
  "strconv"
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
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(books)
}

// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  id := params["id"]
  for _, item := range books {
    if item.ID == id {
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  json.NewEncoder(w).Encode(&Book{})
}

// Create New Book
func createBook(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var book Book
  _ = json.NewDecoder(r.Body).Decode(&book)
  book.ID = strconv.Itoa(rand.Intn(1000000)) // not safe, is repeatable
  books = append(books, book)
  json.NewEncoder(w).Encode(book)
}

// Update Book
func updateBook(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  id := params["id"]
  for i, item := range books {
    if item.ID == id {
      books = append(books[:i], books[i+1:]...)
      var book Book
      _ = json.NewDecoder(r.Body).Decode(&book)
      book.ID = id
      books = append(books, book)
      json.NewEncoder(w).Encode(book)
      return
    }
  }
  json.NewEncoder(w).Encode(books)
}

// Delete Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  id := params["id"]
  for i, item := range books {
    if item.ID == id {
      books = append(books[:i], books[i+1:]...)
      break
    }
  }
  json.NewEncoder(w).Encode(books)
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
      Publisher: &Publisher{Name: "MiddleEarth Binding Co."},
    },
  )
  books = append(
    books, Book{
      ID: "1", Isbn: "110498", Title: "Astrology For You", Price: 22.99,
      Author: &Author{Firstname: "Amanda", Lastname: "Tollefson"},
      Publisher: &Publisher{Name: "Santa Clarita Publishing"},
    },
  )
  books = append(
    books, Book{
      ID: "2", Isbn: "110498", Title: "Lord of the Rings vol 2", Price: 12.99,
      Author: &Author{Firstname: "JRR", Lastname: "Tolkien"},
      Publisher: &Publisher{Name: "MiddleEarth Binding Co."},
    },
  )


  // Route Handlers
  r.HandleFunc("/api/books", getBooks).Methods("GET")
  r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
  r.HandleFunc("/api/books", createBook).Methods("POST")
  r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
  r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":8000", r))
}
