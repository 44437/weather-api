package weather

import (
	"context"
	"fmt"
	"testing"
	mocks "weather-api/internal/weather/mock"
	"weather-api/internal/weather/model"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_GetWeatherByLocation(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("When the service returns an error", func(t *testing.T) {
		mockRepository := mocks.NewMockRepository(mockController)
		mockRepository.
			EXPECT().
			GetWeatherByLocation(gomock.Any(), "ankara").
			Return(
				float32(0),
				fmt.Errorf("error message"),
			).
			Times(1)

		service := NewService(mockRepository)
		response, err := service.GetWeatherByLocation(context.Background(), "ankara")

		assert.Error(t, err)
		assert.Equal(t, &model.Response{}, response)
	})
	t.Run("When the repository returns a successful response", func(t *testing.T) {
		mockRepository := mocks.NewMockRepository(mockController)
		mockRepository.
			EXPECT().
			GetWeatherByLocation(gomock.Any(), "ankara").
			Return(
				float32(23.9),
				nil,
			).
			Times(1)

		service := NewService(mockRepository)
		response, err := service.GetWeatherByLocation(context.Background(), "ankara")

		assert.Nil(t, err)
		assert.Equal(t, &model.Response{
			Location:    "ankara",
			Temperature: float32(23.9),
		}, response)
	})
}
