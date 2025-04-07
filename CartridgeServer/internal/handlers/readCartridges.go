package handlers

import (
	"CartridgeServer/internal/models"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

type Handler struct {
}

type Reader interface {
	ReadCartridges(offset, limit, name string) ([]models.CartridgeModel, error)
}

type ReadResponse struct {
	Cartridges []models.CartridgeModel `json:"cartridges"`
}

const (
	LIMIT_MAX     = 100
	LIMIT_DEFAULT = 20
)

// ReadCartridgesHandler обрабатывает GET запрос по эндпоинту /list
// Генерирует json по ReadResponse
// URL параметры:
//   - limit: default value = 20, max = 100
//   - offset: default value = 0
//   - name: фильтр по названию картриджа/джей
//
// Пример запроса:
// - GET /list?offset=10&limit=5&name=HP
//
// Возможные коды ответа:
// - 200: все ок
// - 400: неверные параметры offset/limit
// - 500: ошибка при чтении данных из бд
func (h *Handler) ReadCartridgesHandler(log *slog.Logger, reader Reader) gin.HandlerFunc {
	return func(c *gin.Context) {

		const op = "CartridgeServer/internal/handlers/ReadCartridgesHandler"
		log.With(slog.String("op", op))

		offset, err := strconv.Atoi(c.Query("offset"))
		if err != nil {
			log.Error("Parsing offset error", err)
			c.JSON(http.StatusBadRequest, ReadResponse{[]models.CartridgeModel{}})
			return
		}
		log.Debug("Parsed offset", offset)

		limit, err := strconv.Atoi(c.Query("offset"))
		if err != nil {
			log.Warn("Parsing limit error", err)
			limit = LIMIT_DEFAULT
		}
		if limit > LIMIT_MAX {
			limit = LIMIT_MAX
		}
		log.Debug("Parsed limit", limit)

		name := c.Query("name")
		log.Debug("Parsed name", name)

	}
}
