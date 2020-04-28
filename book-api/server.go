package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matt-FFFFFF/bookdata-api/datastore"
)

var (
	books datastore.BookStore
)

func init() {
	books = &datastore.Books{}
	books.Initialize()
}

func main() {
	r := mux.NewRouter()
	log.Println("bookdata api")
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "api v1 welcome team 6")
	})
	api.HandleFunc("/books", getAllBooks).Methods(http.MethodGet)
	log.Fatalln(http.ListenAndServe(":8080", r))
}

// https://localhost:8080/api/v1   ===> api v1  / 404

// https://localhost:8080/api/v1/books  ===>  books
