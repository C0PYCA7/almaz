package handlers

import (
	"CartridgeServer/internal/handlers"
	"CartridgeServer/internal/models"
	"CartridgeServer/test/mocks"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReadCartridgesHandler(t *testing.T) {

	mockReader := mocks.NewReader(t)

	expectedCartridges := []models.CartridgeModel{
		{
			Name:          "a",
			Parameters:    "a",
			Status:        "a",
			ReceivedFrom:  "a",
			BarcodeNumber: 228,
		},
		{
			Name:          "b",
			Parameters:    "b",
			Status:        "b",
			ReceivedFrom:  "b",
			BarcodeNumber: 123,
		},
	}

	mockReader.On("ReadCartridges", 0, 20, "").Return(expectedCartridges, nil)

	var log = slog.New(slog.NewTextHandler(&bytes.Buffer{}, nil))

	h := handlers.Handler{}

	r := gin.Default()

	r.GET("/list", h.ReadCartridgesHandler(log, mockReader))

	req, _ := http.NewRequest(http.MethodGet, "/list?offset=1&limit=20", nil)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp handlers.ReadResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, expectedCartridges, resp.Cartridges)

	mockReader.AssertExpectations(t)
}
