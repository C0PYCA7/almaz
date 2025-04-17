package kafka

import (
	"NotificationService/internal/handlers"
	"github.com/IBM/sarama"
	"log"
)

type Consumer struct {
	Ready chan bool
	Hub   *handlers.Hub
}

func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(consumer.Ready)
	return nil
}

func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				log.Printf("message channel was closed")
				return nil
			}
			log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s, partition = %d", string(message.Value), message.Timestamp, message.Topic, message.Partition)

			if consumer.Hub != nil {
				consumer.Hub.Broadcast <- message.Value
			}
			session.MarkMessage(message, "")

		case <-session.Context().Done():
			log.Println("Session context done, rebalance might be happening")
			return nil
		}
	}
}
