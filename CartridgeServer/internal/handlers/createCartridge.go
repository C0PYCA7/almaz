package handlers

import (
	"CartridgeServer/internal/models"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type CreateRequest struct {
	Cartridges []models.CreateCartridgeModel `json:"cartridges"`
}

type CreateResponse struct {
	Message string `json:"message"`
}

func (h *Handler) CreateCartridgeHandler(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		var (
			messages []models.DbTopicMessage
			message  models.DbTopicMessage
		)

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
				Action: "create",
				Cartridge: models.CartridgeModel{
					Name:          cartridge.Name,
					Parameters:    cartridge.Parameters,
					ReceivedFrom:  cartridge.ReceivedFrom,
					BarcodeNumber: cartridge.BarcodeNumber,
				},
			}
			messages = append(messages, message)
		}
	}
}
