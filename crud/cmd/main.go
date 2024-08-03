package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"crud/internal/movies"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/movies", movies.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", movies.GetMovie).Methods("GET")
	r.HandleFunc("/movies", movies.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", movies.UpdateMoive).Methods("PUT")
	r.HandleFunc("/movies/{id}", movies.DeleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 3031")

	if err := http.ListenAndServe(":3031", r); err != nil {
		log.Fatal(err)
	}
}
