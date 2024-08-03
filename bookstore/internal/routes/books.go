package routes

import (
	"bookstore/internal/handlers"

	"github.com/gorilla/mux"
)

var BookstoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	router.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", handlers.GetBookByID).Methods("GET")
	router.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")

}
