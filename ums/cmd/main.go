package main

import (
	"net/http"
	"ums/internal/config"
	"ums/internal/models"

	"github.com/julienschmidt/httprouter"

	"ums/internal/handlers"
)

func main() {
	r := httprouter.New()

	s := config.GetSession()
	defer s.Close()

	svc := models.NewSvc(s)
	uh := handlers.NewUserHandler(svc)

	r.POST("/users", uh.CreateUser)
	r.GET("/usesr/:id", uh.GetUserByID)
	r.PUT("/users/:id", uh.UpdateUser)
	r.DELETE("/users/:id", uh.DeleteUser)

	http.ListenAndServe("localhost:9000", r)
}
