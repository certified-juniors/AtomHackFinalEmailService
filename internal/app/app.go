package app

import (
	"github.com/certified-juniors/AtomHackFinalEmailService/internal/config"
	"github.com/certified-juniors/AtomHackFinalEmailService/internal/http/handler"
	"github.com/certified-juniors/AtomHackFinalEmailService/internal/smtp"
)

type Application struct {
	cfg     *config.App
	handler *handler.Handler
}

func New() (*Application, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	smtp := smtp.NewSMTP(&cfg.SMTP)

	h := handler.New(smtp)

	app := &Application{
		cfg:     cfg,
		handler: h,
	}

	return app, nil
}
