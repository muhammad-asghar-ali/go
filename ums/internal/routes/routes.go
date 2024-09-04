package routes

import (
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"

	"ums/internal/handlers"
	"ums/internal/models"
)

func RegisterRoutes(router *httprouter.Router, client *mongo.Client) {
	userService := models.NewSvc(client)
	userHandler := handlers.NewUserHandler(userService)

	router.POST("/users", userHandler.CreateUser)
	router.GET("/users/:id", userHandler.GetUserByID)
	router.PUT("/users/:id", userHandler.UpdateUser)
	router.DELETE("/users/:id", userHandler.DeleteUser)
}
