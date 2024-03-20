package config

type App struct {
	ErrorLevel string `env:"EMAIL_SERVICE_ERROR_LEVEL" envDefault:"info"`

	API   API
	SMTP  SMTP
	Kafka Kafka
}
