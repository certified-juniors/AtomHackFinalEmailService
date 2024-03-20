package handler

import (
	"github.com/certified-juniors/AtomHackFinalEmailService/internal/smtp"
)

type Handler struct {
	s *smtp.SMTP
}

func New(SMTP *smtp.SMTP) *Handler {
	return &Handler{
		s: SMTP, 
	}
}
