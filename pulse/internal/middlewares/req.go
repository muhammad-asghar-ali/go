package middlewares

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"pulse/internal/errors"
)

func GetUserIDFromRequest(ctx *gin.Context) (string, error) {
	userID := ctx.Param("userID")
	if userID == "" {
		return "", errors.ErrNoMessagesFound
	}

	return userID, nil
}

func GetIDFromRequest(formValue string, ctx *gin.Context) (int, error) {
	id, err := strconv.Atoi(ctx.PostForm(formValue))
	if err != nil {
		return 0, fmt.Errorf(
			"failed to parse ID from form value %s: %w", formValue, err)
	}

	return id, nil
}
