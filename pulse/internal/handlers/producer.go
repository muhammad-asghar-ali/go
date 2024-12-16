package handlers

import (
	er "errors"
	"net/http"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"

	"pulse/internal/errors"
	"pulse/internal/kafka"
	"pulse/internal/middlewares"
	"pulse/internal/models"
)

func SendMessageHandler(producer sarama.SyncProducer, users []models.User) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fromID, err := middlewares.GetIDFromRequest("fromID", ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		toID, err := middlewares.GetIDFromRequest("toID", ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		err = kafka.SendKafkaMessage(producer, users, ctx, fromID, toID)
		if er.Is(err, errors.ErrUserNotFoundInProducer) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		}

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Notification sent successfully!",
		})
	}
}
