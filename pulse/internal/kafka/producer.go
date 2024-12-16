package kafka

import (
	"encoding/json"
	"fmt"
	"pulse/internal/models"
	"pulse/internal/utils"
	"strconv"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
)

func SetupProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{utils.KafkaServerAddress}, config)
	if err != nil {
		return nil, fmt.Errorf("failed to setup producer: %w", err)
	}

	return producer, nil
}

func SendKafkaMessage(
	producer sarama.SyncProducer, users []models.User, ctx *gin.Context, fromID, toID int,
) error {
	message := ctx.PostForm("message")
	u := models.User{}

	fromUser, err := u.FindUserByID(fromID, users)
	if err != nil {
		return err
	}

	toUser, err := u.FindUserByID(toID, users)
	if err != nil {
		return err
	}

	notification := models.Notification{
		From: fromUser,
		To:   toUser, Message: message,
	}

	notificationJSON, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("failed to marshal notification: %w", err)
	}

	msg := &sarama.ProducerMessage{
		Topic: utils.KafkaTopic,
		Key:   sarama.StringEncoder(strconv.Itoa(toUser.ID)),
		Value: sarama.StringEncoder(notificationJSON),
	}

	_, _, err = producer.SendMessage(msg)

	return err
}
