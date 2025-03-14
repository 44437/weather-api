package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	envProd = "prod"
)

type Handler interface {
	RegisterRoutes(app *echo.Echo)
}

type Server interface {
	GetEchoInstance() *echo.Echo
	Start() error
	Stop()
}

type server struct {
	echo *echo.Echo
}

func NewServer(port int, handlers []Handler) Server {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.Server.Addr = fmt.Sprintf(":%d", port)

	server := &server{
		echo: e,
	}
	server.addRoutes()

	for _, handler := range handlers {
		handler.RegisterRoutes(server.echo)
	}

	return server
}

func (s *server) addRoutes() {
	s.echo.GET("/health", healthCheck)

	if os.Getenv("APP_ENV") != envProd {
		s.echo.GET("/debug/pprof/*", echo.WrapHandler(http.DefaultServeMux))
	}
}

func (s *server) GetEchoInstance() *echo.Echo {
	return s.echo
}

func healthCheck(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func (s *server) Start() error {
	return s.echo.Server.ListenAndServe()
}

func (s *server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := s.echo.Server.Shutdown(ctx)
	if err != nil {
		log.Printf("Server shutdown failed: %v", err)
	}
}
