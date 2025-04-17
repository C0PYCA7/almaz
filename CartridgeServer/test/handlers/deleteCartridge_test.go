package handlers

//func TestDeleteCartridgeHandler(t *testing.T) {
//
//	var testCases = []struct {
//		name          string
//		inputQuery    string
//		mockInput     int
//		mockReturnErr error
//		expectedCode  int
//		expectedBody  string
//	}{
//		{
//			name:          "valid delete",
//			inputQuery:    "number=228357",
//			mockInput:     228357,
//			mockReturnErr: nil,
//			expectedCode:  http.StatusOK,
//			expectedBody:  "Cartridge deleted",
//		},
//		{
//			name:         "invalid number",
//			inputQuery:   "number=invalid",
//			expectedCode: http.StatusBadRequest,
//			expectedBody: "failed to parse barcodeNumber",
//		},
//		{
//			name:          "cartridge not found",
//			inputQuery:    "number=999999",
//			mockInput:     999999,
//			mockReturnErr: pgx.ErrNoRows,
//			expectedCode:  http.StatusNotFound,
//			expectedBody:  "Cartridge does not exist",
//		},
//		{
//			name:          "internal error",
//			inputQuery:    "number=500",
//			mockInput:     500,
//			mockReturnErr: errors.New("some db error"),
//			expectedCode:  http.StatusInternalServerError,
//			expectedBody:  "Failed to delete cartridge",
//		},
//	}
//
//	for _, testCase := range testCases {
//		t.Run(testCase.name, func(t *testing.T) {
//			mockDeleter := mocks.NewDeleter(t)
//
//			if testCase.expectedCode != http.StatusBadRequest {
//				mockDeleter.On("DeleteCartridge", testCase.mockInput).Return(testCase.mockReturnErr)
//			}
//
//			log := slog.New(slog.NewTextHandler(&bytes.Buffer{}, nil))
//			h := handlers.Handler{}
//			r := gin.Default()
//			r.DELETE("/delete", h.DeleteCartridgeHandler(log, mockDeleter))
//
//			req, _ := http.NewRequest(http.MethodDelete, "/delete?"+testCase.inputQuery, nil)
//			w := httptest.NewRecorder()
//			r.ServeHTTP(w, req)
//
//			assert.Equal(t, testCase.expectedCode, w.Code)
//			assert.Contains(t, w.Body.String(), testCase.expectedBody)
//		})
//	}
//
//}
