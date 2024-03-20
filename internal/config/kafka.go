package config

type Kafka struct {
	Addr            string `env:"EMAIL_SERVICE_BOOTSTRAP_SERVER"`
	Topic           string `env:"EEMAIL_SERVICE_KAFKA_TOPIC"`
	MaxRetry        int    `env:"EMAIL_SERVICE_MAX_RETRY"`
	ReturnSuccesses bool   `env:"EMAIL_SERVICE_RETURN_SUCCESSES"`
}
