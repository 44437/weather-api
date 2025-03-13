package main

import (
	"log"
	"weather-api/internal/server"

	"go.uber.org/zap"
)

func main() {
	srv := server.NewServer()

	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)

	err := srv.Start()
	if err != nil {
		log.Fatalln(err)
	}
}
