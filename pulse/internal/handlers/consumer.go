package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"pulse/internal/middlewares"
	"pulse/internal/models"
	"pulse/internal/store"
)

func HandleNotifications(ctx *gin.Context, store *store.NotificationStore) {
	userID, err := middlewares.GetUserIDFromRequest(ctx)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	notes := store.Get(userID)
	if len(notes) == 0 {
		ctx.JSON(http.StatusOK,
			gin.H{
				"message":       "No notifications found for user",
				"notifications": []models.Notification{},
			})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"notifications": notes})
}
