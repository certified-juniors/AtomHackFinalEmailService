package config

type API struct {
	ServiceHost string `env:"EMAIL_SERVICE_HOST" envDefault:"0.0.0.0"`
	ServicePort int    `env:"EMAIL_SERVICE_PORT" envDefault:"8081"`
}
