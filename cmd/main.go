package main

import (
	"log"

	"github.com/certified-juniors/AtomHackFinalEmailService/internal/app"
)

// @title AtomHackFinalEmailService RestAPI
// @version 1.0
// @description API server for EmailService application

// @host http://localhost:8081
// @BasePath /api

func main() {
	log.Println("Application start!")

	application, err := app.New()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	application.Run()
	log.Println("Application terminated!")
}
