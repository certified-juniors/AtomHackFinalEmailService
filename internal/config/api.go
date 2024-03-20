package config

type API struct {
	ServiceHost string `env:"EMAIL_SERVICE_HOST"`
	ServicePort int    `env:"EMAIL_SERVICE_PORT"`
}
