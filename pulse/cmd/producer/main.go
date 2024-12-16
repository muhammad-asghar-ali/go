package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"pulse/internal/handlers"
	"pulse/internal/kafka"
	"pulse/internal/models"
	"pulse/internal/utils"
)

func main() {
	users := []models.User{
		{ID: 1, Name: "John"},
		{ID: 2, Name: "Ali"},
		{ID: 3, Name: "Jana"},
		{ID: 4, Name: "Lena"},
	}

	producer, err := kafka.SetupProducer()
	if err != nil {
		log.Fatalf("failed to initialize producer: %v", err)
	}
	defer producer.Close()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/send", handlers.SendMessageHandler(producer, users))

	fmt.Printf("Kafka PRODUCER 📨 started at http://localhost%s\n", utils.ProducerPort)

	if err := router.Run(utils.ProducerPort); err != nil {
		log.Printf("failed to run the server: %v", err)
	}
}
