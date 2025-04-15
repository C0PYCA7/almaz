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

type UpdateReceiveRequest struct {
	BarcodeNumber int    `json:"barcodeNumber"`
	NewStatus     string `json:"newStatus"`
	//ReceiveFrom   string `json:"receiveFrom"` // точно ли он нужен?
}

type UpdateReceiveResponse struct {
	BarcodeNumber int    `json:"barcodeNumber"`
	Message       string `json:"message"`
}

func (h *Handler) UpdateReceiveCartridgeHandler(log *slog.Logger, sender kafka.Sender, config config.KafkaConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = "CartridgeServer/internal/handlers/UpdateReceiveCartridge"

		log = log.With("op", op)
		var req UpdateReceiveRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			log.Error("failed to decode request", slog.Any("err", err))
			c.JSON(http.StatusBadRequest, UpdateReceiveResponse{
				BarcodeNumber: req.BarcodeNumber,
				Message:       "failed to decode request",
			})
			return
		}

		message := models.DbTopicMessage{
			Action:        "updateReceive",
			BarcodeNumber: req.BarcodeNumber,
			NewStatus:     req.NewStatus,
			Timestamp:     time.Now(),
			//ReceivedFrom:  req.ReceiveFrom,
		}

		dataBytes, err := json.Marshal(message)
		if err != nil {
			log.Error("failed to encode message", slog.Any("err", err))
			c.JSON(http.StatusBadRequest, UpdateReceiveResponse{
				BarcodeNumber: req.BarcodeNumber,
				Message:       "failed to encode message",
			})
			return
		}
		sender.SendMessage(config.DbTopic, dataBytes)
		c.JSON(http.StatusOK, UpdateReceiveResponse{
			BarcodeNumber: req.BarcodeNumber,
			Message:       "Сообщения приняты, вскоре будут обработаны",
		})
		return
	}
}
