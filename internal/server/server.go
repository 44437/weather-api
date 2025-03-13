package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handler interface {
	RegisterRoutes(app *echo.Echo)
}

type Server interface {
	Start() error
	Stop()
}

type server struct {
	echo *echo.Echo
}

func NewServer() Server {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.Server.Addr = fmt.Sprintf(":%d", 8080)

	server := &server{
		echo: e,
	}
	server.addRoutes()

	return server
}

func (s *server) addRoutes() {
	s.echo.GET("/health", healthCheck)
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
		log.Fatalf("server shutdown failed: %+s", err)
	}
}
