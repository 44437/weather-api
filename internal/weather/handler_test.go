package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	mocks "weather-api/internal/weather/mock"
	"weather-api/internal/weather/model"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestRegisterRoutes(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockService := mocks.NewMockService(mockController)
	h := NewHandler(mockService)
	e := echo.New()
	h.RegisterRoutes(e)
	routes := e.Routes()

	assert.Equal(t, "GET", routes[0].Method)
	assert.Equal(t, "/weather", routes[0].Path)
}

func getMockRequest(location string) (echo.Context, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/weather?q=%s", location), nil)
	req.Header.Set("Accept", "application/json")

	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)

	return c, rec
}

func TestHandler_GetWeatherByLocation(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("When the location is empty", func(t *testing.T) {
		ctx, recorder := getMockRequest("")

		handler := NewHandler(nil)
		_ = handler.GetWeatherByLocation(ctx)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		assert.Equal(t, 0, recorder.Body.Len())
	})

	t.Run("When the service returns an error", func(t *testing.T) {
		ctx, recorder := getMockRequest("ankara")

		mockService := mocks.NewMockService(mockController)
		mockService.
			EXPECT().
			GetWeatherByLocation(gomock.Any(), "ankara").
			Return(
				&model.Response{},
				fmt.Errorf("error message"),
			).
			Times(1)

		handler := NewHandler(mockService)
		_ = handler.GetWeatherByLocation(ctx)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
		assert.Equal(t, "error message", recorder.Body.String())
	})
	t.Run("When the handler returns a successful response", func(t *testing.T) {
		ctx, recorder := getMockRequest("ankara")

		mockService := mocks.NewMockService(mockController)
		mockService.
			EXPECT().
			GetWeatherByLocation(gomock.Any(), "ankara").
			Return(
				&model.Response{
					Location:    "ankara",
					Temperature: 23.9,
				},
				nil,
			).
			Times(1)

		handler := NewHandler(mockService)
		_ = handler.GetWeatherByLocation(ctx)

		assert.Equal(t, http.StatusOK, recorder.Code)

		var response model.Response
		_ = json.Unmarshal(recorder.Body.Bytes(), &response)
		assert.Equal(t, "ankara", response.Location)
		assert.Equal(t, float32(23.9), response.Temperature)
	})
}
