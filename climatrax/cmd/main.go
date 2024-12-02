package main

import (
	"fmt"
	"log"
	"net/http"

	"climatrax/handlers"
)

func main() {
	http.HandleFunc("/hello", handlers.Hello)
	http.HandleFunc("/weather/", handlers.Weather)

	port := ":8080"
	fmt.Println("Starting server on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
