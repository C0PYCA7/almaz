package handlers

import (
	"CartridgeServer/internal/config"
	"CartridgeServer/internal/kafka"
	"CartridgeServer/internal/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"time"
)

type CreateRequest struct {
	Cartridges []models.CreateCartridgeModel `json:"cartridges"`
}

type CreateResponse struct {
	Message string `json:"message"`
}

const RECEIVED_FROM = "получено из подразделения"

func (h *Handler) CreateCartridgeHandler(log *slog.Logger, sender kafka.Sender, config config.KafkaConfig) gin.HandlerFunc {
	return func(c *gin.Context) {

		var message models.DbTopicMessage

		const op = "CartridgeServer/internal/handlers/CreateCartridge"
		log = log.With("op", op)

		var req CreateRequest

		if err := c.BindJSON(&req); err != nil {
			{
				log.Error("failed to decode body", slog.Any("err", err))
				c.JSON(http.StatusBadRequest, CreateResponse{
					Message: "failed to decode body",
				})
				return
			}
		}
		for _, cartridge := range req.Cartridges {
			message = models.DbTopicMessage{
				Action:        "create",
				Name:          cartridge.Name,
				Parameters:    cartridge.Parameters,
				BarcodeNumber: cartridge.BarcodeNumber,
				Timestamp:     time.Now(),
				ReceivedFrom:  cartridge.ReceivedFrom,
				NewStatus:     RECEIVED_FROM,
			}
			dataBytes, err := json.Marshal(message)
			if err != nil {
				log.Error("failed to encode message", slog.Any("err", err))
				c.JSON(http.StatusBadRequest, CreateResponse{
					Message: "failed to encode message",
				})
				return
			}
			sender.SendMessage(config.DbTopic, dataBytes)
		}
		c.JSON(http.StatusOK, CreateResponse{
			Message: "Сообщения приняты, вскоре будут обработаны",
		})
		return
	}
}
