package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"pulse/internal/models"
	"pulse/internal/store"
	"pulse/internal/utils"

	"github.com/IBM/sarama"
)

type (
	Consumer struct {
		store *store.NotificationStore
	}
)

func (*Consumer) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (*Consumer) Cleanup(sarama.ConsumerGroupSession) error { return nil }

func (consumer *Consumer) ConsumeClaim(
	sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	for msg := range claim.Messages() {
		userID := string(msg.Key)
		var notification models.Notification
		err := json.Unmarshal(msg.Value, &notification)
		if err != nil {
			log.Printf("failed to unmarshal notification: %v", err)
			continue
		}
		consumer.store.Add(userID, notification)

		sess.MarkMessage(msg, "")
	}

	return nil
}

func InitializeConsumerGroup() (sarama.ConsumerGroup, error) {
	config := sarama.NewConfig()

	consumerGroup, err := sarama.NewConsumerGroup([]string{utils.KafkaServerAddress}, utils.ConsumerGroup, config)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize consumer group: %w", err)
	}

	return consumerGroup, nil
}

func SetupConsumerGroup(ctx context.Context, store *store.NotificationStore) {
	consumerGroup, err := InitializeConsumerGroup()
	if err != nil {
		log.Printf("initialization error: %v", err)
	}
	defer consumerGroup.Close()

	consumer := &Consumer{
		store: store,
	}

	for {
		if err := consumerGroup.Consume(ctx, []string{utils.ConsumerTopic}, consumer); err != nil {
			log.Printf("error from consumer: %v", err)
		}

		if ctx.Err() != nil {
			return
		}
	}
}
