package handlers

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

// TODO: поменять response, потому что данные будут идти в кафку и ответа ждать такого не стоит
type DeleteResponse struct {
	Message string `json:"message"`
}

/*
DeleteCartridgeHandler обрабатывает DELETE запрос по эндпоинду /delete/{number}
TODO: поменять на вызов кафки
*/
func (h *Handler) DeleteCartridgeHandler(log *slog.Logger) gin.HandlerFunc {
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

		//err = deleter.DeleteCartridge(number)
		//if err != nil {
		//	if errors.Is(err, pgx.ErrNoRows) {
		//		log.Error("cartridge does not exist", slog.Any("error", err))
		//		c.JSON(http.StatusNotFound, DeleteResponse{
		//			BarcodeNumber: number,
		//			Message:       "Cartridge does not exist",
		//		})
		//		return
		//	}
		//	log.Error("Deleting cartridge error", slog.Any("error", err))
		//	c.JSON(http.StatusInternalServerError, DeleteResponse{
		//		BarcodeNumber: number,
		//		Message:       "Failed to delete cartridge",
		//	})
		//	return
		//}
		//
		//log.Info("Cartridge deleted", slog.String("number", numberString))
		//c.JSON(http.StatusOK, DeleteResponse{
		//	BarcodeNumber: number,
		//	Message:       "Cartridge deleted",
		//})
		//return
	}
}
