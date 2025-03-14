package main

import (
	"log"
	"os"
	"weather-api/internal/config"
	"weather-api/internal/server"
	"weather-api/internal/weather"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)

	appEnv := os.Getenv("APP_ENV")
	conf, err := config.New(".config", appEnv)
	if err != nil {
		log.Fatalln(err)
	}
	conf.Print()

	handler := weather.NewHandler(weather.NewService(weather.NewRepository(conf.Postgres)))

	handlers := []server.Handler{handler}

	srv := server.NewServer(conf.Server.Port, handlers)

	err = srv.Start()
	if err != nil {
		log.Fatalln(err)
	}
}
