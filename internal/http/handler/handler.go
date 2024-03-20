package handler

import (
	"github.com/certified-juniors/AtomHackFinalEmailService/internal/smtp"
)

type Handler struct {
	// p *kafka.Producer
	s *smtp.SMTP
}

func New(SMTP *smtp.SMTP) *Handler {
	// producer, err := kafka.NewProducer(&config.Kafka)
	// if err != nil {
	// 	log.Fatal("Error occured while creating producer", err)
	// }
	return &Handler{
		// p: producer,
		s: SMTP, 
	}
}
