package weather

import (
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	RegisterRoutes(e *echo.Echo)
	GetWeatherByLocation(ec echo.Context) error
}

type handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &handler{service: service}
}

func (h *handler) RegisterRoutes(e *echo.Echo) {
	e.GET("/weather", h.GetWeatherByLocation)
}

func (h *handler) GetWeatherByLocation(ctx echo.Context) error {
	location := ctx.QueryParam("q")

	if location == "" {
		log.Println("Location cannot be empty for this request") //TODO middleware?

		return ctx.NoContent(http.StatusBadRequest)
	}

	location = strings.TrimSpace(location)
	location = strings.ToLower(location)

	response, err := h.service.GetWeatherByLocation(ctx.Request().Context(), location)
	if err != nil {
		log.Printf("There was a problem in handler#GetWeatherByLocation: %v", err)

		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, *response)
}
