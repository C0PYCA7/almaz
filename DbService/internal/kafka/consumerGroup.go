package kafka

import (
	"DbService/internal/models"
	"DbService/internal/storage"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/IBM/sarama"
	"log"
)

type Database interface {
	CreateCartridge(cartridge *models.CreateCartridge) error
	UpdateCartridgeReceiveStatus(cartridge *models.UpdateCartridgeReceive) error
	UpdateCartridgeSendStatus(cartridge *models.UpdateCartridgeSend) error
	DeleteCartridge(barcodeNumber int) error
}

const (
	ACTION_CREATE         = "create"
	ACTION_UPDATE_RECEIVE = "updateReceive"
	ACTION_UPDATE_SEND    = "updateSend"
	ACTION_DELETE         = "delete"
)

type Consumer struct {
	Ready    chan bool
	Database Database
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
			var msg models.DbTopicMessage
			err := json.Unmarshal(message.Value, &msg)
			if err != nil {
				log.Println("unmarshal message in consume: ", err)
				return err
			}

			log.Println(msg)

			err = consumer.handleMessage(&msg)
			if err != nil {
				if errors.Is(err, storage.ErrNotFound) {
					log.Println("message not found: ", err)
					session.MarkMessage(message, "")
				}
				if errors.Is(err, storage.ErrUniqueBarcode) {
					log.Println("barcode already exists: ", err)
					session.MarkMessage(message, "")
				}
				log.Println("message handle message failed: ", err)
				continue
			}

			session.MarkMessage(message, "")

		case <-session.Context().Done():
			log.Println("Session context done, rebalance might be happening")
			return nil
		}
	}
}

func (consumer *Consumer) handleMessage(msg *models.DbTopicMessage) error {
	switch msg.Action {
	case ACTION_CREATE:
		cartridge := &models.CreateCartridge{
			Name:          msg.Name,
			Parameters:    msg.Parameters,
			BarcodeNumber: msg.BarcodeNumber,
			ReceivedFrom:  msg.ReceivedFrom,
			Timestamp:     msg.Timestamp,
		}
		return consumer.Database.CreateCartridge(cartridge)

	case ACTION_UPDATE_SEND:
		update := &models.UpdateCartridgeSend{
			BarcodeNumber: msg.BarcodeNumber,
			NewStatus:     msg.NewStatus,
			SendTo:        msg.SendTo,
			Timestamp:     msg.Timestamp,
		}
		return consumer.Database.UpdateCartridgeSendStatus(update)

	case ACTION_UPDATE_RECEIVE:
		update := &models.UpdateCartridgeReceive{
			BarcodeNumber: msg.BarcodeNumber,
			NewStatus:     msg.NewStatus,
			Timestamp:     msg.Timestamp,
		}
		return consumer.Database.UpdateCartridgeReceiveStatus(update)

	case ACTION_DELETE:
		return consumer.Database.DeleteCartridge(msg.BarcodeNumber)

	default:
		return fmt.Errorf("unknown action: %s", msg.Action)
	}
}
