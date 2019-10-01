package main

import (
  "log"
  "net/http"
  "math/rand"
  "strconv"
  "encoding/json"
  "github.com/gorilla/mux"
)

type Book struct {
  Id string `json:"id"`
  Title string `json:"title"`
  Author *Author `json:"author"`
}

type Author struct {
  Firstname string `json:"firstname"`
  Lastname string `json:"lastname"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for _, item := range books {
     if item.Id == params["id"] {
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var book Book
  _ = json.NewDecoder(r.Body).Decode(&book)
  book.Id = strconv.Itoa(rand.Intn(1000000))
  books = append(books, book)
  json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for index, item := range books {
    if item.Id == params["id"] {
      books = append(books[:index], books[index+1:]...)
      var book Book
      _ = json.NewDecoder(r.Body).Decode(&book)
      book.Id = params["id"]
      books = append(books, book)
      json.NewEncoder(w).Encode(book)
      return
    }
  }
  json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for index, item := range books {
    if item.Id == params["id"] {
      books = append(books[:index], books[index+1:]...)
      break
    }
  }
  json.NewEncoder(w).Encode(books)
}

func main() {
  r := mux.NewRouter()
  books = append(books, Book{Id: "1", Title: "Война и Мир", Author: &Author{Firstname: "Лев", Lastname: "Толстой"}})
  books = append(books, Book{Id: "2", Title: "Преступление и наказание", Author: &Author{Firstname: "Фёдор", Lastname: "Достоевский"}})
  r.HandleFunc("/books", getBooks).Methods("GET")
  r.HandleFunc("/book/{id}", getBook).Methods("GET")
  r.HandleFunc("/book", createBook).Methods("POST")
  r.HandleFunc("/book/{id}", updateBook).Methods("PUT")
  r.HandleFunc("/book/{id}", deleteBook).Methods("DELETE")
  log.Fatal(http.ListenAndServe(":8080", r))
}
