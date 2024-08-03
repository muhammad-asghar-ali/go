package main

import (
	"bookstore/internal/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.BookstoreRoutes(r)
	http.Handle("/", r)

	fmt.Println("Starting server at port 3032")

	if err := http.ListenAndServe(":3032", r); err != nil {
		log.Fatal(err)
	}
}
