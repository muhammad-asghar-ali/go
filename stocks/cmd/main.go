package main

import (
	"fmt"
	"log"
	"net/http"

	"stocks/internal/routes"
)

func main() {
	r := routes.Router()

	fmt.Println("Starting server at port 8082")

	if err := http.ListenAndServe(":8082", r); err != nil {
		log.Fatal(err)
	}
}
