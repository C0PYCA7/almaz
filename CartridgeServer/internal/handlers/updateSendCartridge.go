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

type UpdateSendRequest struct {
	BarcodeNumber int    `json:"barcodeNumber"`
	NewStatus     string `json:"newStatus"`
	SendTo        string `json:"sendTo"`
}

type UpdateSendResponse struct {
	BarcodeNumber int    `json:"barcodeNumber"`
	Message       string `json:"message"`
}

func (h *Handler) UpdateSendCartridgeHandler(log *slog.Logger, sender kafka.Sender, config config.KafkaConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = "CartridgeServer/internal/handlers/UpdateSendCartridge"

		log = log.With("op", op)
		var req UpdateSendRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			log.Error("failed to decode request", slog.Any("err", err))
			c.JSON(http.StatusBadRequest, UpdateSendResponse{
				BarcodeNumber: req.BarcodeNumber,
				Message:       "failed to decode request",
			})
			return
		}

		message := models.DbTopicMessage{
			Action:        "update",
			BarcodeNumber: req.BarcodeNumber,
			NewStatus:     req.NewStatus,
			Timestamp:     time.Now(),
			SendTo:        req.SendTo,
		}

		dataBytes, err := json.Marshal(message)
		if err != nil {
			log.Error("failed to encode message", slog.Any("err", err))
			c.JSON(http.StatusBadRequest, UpdateReceiveResponse{
				Message: "failed to encode message",
			})
			return
		}
		sender.SendMessage(config.DbTopic, dataBytes)
		c.JSON(http.StatusOK, UpdateSendResponse{
			Message: "Сообщения приняты, вскоре будут обработаны",
		})
		return
	}
}
