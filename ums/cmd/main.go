package main

import (
	"log"
	"net/http"
	"ums/internal/config"
	"ums/internal/models"

	"github.com/julienschmidt/httprouter"

	"ums/internal/handlers"
)

func main() {
	r := httprouter.New()

	s := config.GetSession()
	if s == nil {
		log.Fatal("Could not establish a MongoDB session")
	}
	defer s.Close()

	svc := models.NewSvc(s)
	uh := handlers.NewUserHandler(svc)

	r.POST("/users", uh.CreateUser)
	r.GET("/users/:id", uh.GetUserByID)
	r.PUT("/users/:id", uh.UpdateUser)
	r.DELETE("/users/:id", uh.DeleteUser)

	log.Fatal(http.ListenAndServe(":9000", r))
}
