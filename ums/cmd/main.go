package main

import (
	"net/http"
	"ums/internal/config"

	"github.com/julienschmidt/httprouter"

	"ums/internal/handlers"
)

func main() {
	r := httprouter.New()

	uh := handlers.NewUserHandler(config.GetSession())

	r.POST("/users", uh.CreateUser)
	r.GET("/usesr/:id", uh.GetUserByID)
	r.PUT("/users/:id", uh.UpdateUser)
	r.DELETE("/users/:id", uh.DeleteUser)

	http.ListenAndServe("localhost:9000", r)
}
