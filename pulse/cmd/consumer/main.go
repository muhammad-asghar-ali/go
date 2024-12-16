package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"pulse/internal/handlers"
	"pulse/internal/kafka"
	"pulse/internal/store"
	"pulse/internal/utils"
)

func main() {
	store := &store.NotificationStore{
		Data: make(store.UserNotifications),
	}

	ctx, cancel := context.WithCancel(context.Background())
	go kafka.SetupConsumerGroup(ctx, store)
	defer cancel()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/notifications/:userID", func(ctx *gin.Context) {
		handlers.HandleNotifications(ctx, store)
	})

	fmt.Printf("Kafka CONSUMER (Group: %s) 👥📥 "+
		"started at http://localhost%s\n", utils.ConsumerGroup, utils.ConsumerPort)

	if err := router.Run(utils.ConsumerPort); err != nil {
		log.Printf("failed to run the server: %v", err)
	}
}
