package main

import (
	"context"
	"log"
	"net/http"
	"ums/internal/config"
	"ums/internal/routes"

	"github.com/julienschmidt/httprouter"
)

func main() {
	client := config.InitDB()
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	router := httprouter.New()

	routes.RegisterRoutes(router, client)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
