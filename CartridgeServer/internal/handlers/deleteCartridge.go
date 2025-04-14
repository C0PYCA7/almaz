package handlers

import (
	"CartridgeServer/internal/config"
	"CartridgeServer/internal/kafka"
	"CartridgeServer/internal/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

type DeleteResponse struct {
	Message string `json:"message"`
}

/*
DeleteCartridgeHandler обрабатывает DELETE запрос по эндпоинду /delete/{number}
*/
func (h *Handler) DeleteCartridgeHandler(log *slog.Logger, sender kafka.Sender, config config.KafkaConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = "CartridgeServer/internal/handlers/DeleteCartridgeHandler"
		log = log.With(slog.String("op", op))

		numberString := c.Query("number")
		log.Debug("Receive number", slog.String("number", numberString))

		number, err := strconv.Atoi(numberString)
		if err != nil {
			log.Error("Parsing barcodeNumber error", slog.Any("error", err))
			c.JSON(http.StatusBadRequest, DeleteResponse{
				Message: "failed to parse barcodeNumber",
			})
			return
		}
		log.Debug("Parsed number", slog.Int("number", number))

		message := models.DbTopicMessage{
			Action:        "delete",
			BarcodeNumber: number,
		}

		dataBytes, err := json.Marshal(message)
		if err != nil {
			log.Error("failed to encode message", slog.Any("err", err))
			c.JSON(http.StatusInternalServerError, DeleteResponse{
				Message: "failed to encode message",
			})
			return
		}
		sender.SendMessage(config.DbTopic, dataBytes)

		c.JSON(http.StatusOK, DeleteResponse{
			Message: "Сообщения приняты, вскоре будут обработаны",
		})
		return
	}
}
