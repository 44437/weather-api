package weather

import (
	"context"
	"weather-api/internal/weather/model"
)

type Service interface {
	GetWeatherByLocation(ctx context.Context, location string) (*model.Response, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) GetWeatherByLocation(ctx context.Context, location string) (*model.Response, error) {
	averageTemperature, err := s.repository.GetWeatherByLocation(ctx, location)
	if err != nil {
		return &model.Response{}, err
	}

	return &model.Response{
		Location:    location,
		Temperature: averageTemperature,
	}, nil
}
