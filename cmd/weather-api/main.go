package main

import (
	"log"
	"weather-api/internal/server"
)

func main() {
	srv := server.NewServer()
	err := srv.Start()

	if err != nil {
		log.Fatalln(err)
	}
}
