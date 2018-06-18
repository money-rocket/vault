package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"encoding/json"
	"math/rand"
	"strconv"
)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var books []Book

func getBooks(resp http.ResponseWriter, request *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(books);
}

func getBook(resp http.ResponseWriter, request *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(resp).Encode(item)
			return
		}
	}

	json.NewEncoder(resp).Encode(&Book{})
}

func createBook(resp http.ResponseWriter, request *http.Request) {
	resp.Header().Set("Content-Type", "application/json")

	var book Book
	_ = json.NewDecoder(request.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000))

	books = append(books, book)
	json.NewEncoder(resp).Encode(book);
}

func updateBook(resp http.ResponseWriter, request *http.Request) {

}

func deleteBook(resp http.ResponseWriter, request *http.Request) {

}

func main() {
	router := mux.NewRouter()

	books = append(books, Book{ID: "1", Isbn: "12345", Title: "Book one", Author: &Author{FirstName: "Lorem", LastName: "Dolors"}})
	books = append(books, Book{ID: "2", Isbn: "6789", Title: "Book two", Author: &Author{FirstName: "Dolor", LastName: "Dolors"}})
	books = append(books, Book{ID: "3", Isbn: "101112", Title: "Book three", Author: &Author{FirstName: "Elere", LastName: "Dolors"}})
	books = append(books, Book{ID: "4", Isbn: "131415", Title: "Book four", Author: &Author{FirstName: "Ipsun", LastName: "Dolors"}})

	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router));
}
