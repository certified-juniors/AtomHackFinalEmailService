package config

type SMTP struct {
	SmtpHost        string `env:"MAIL_SERVICE_SMTP_HOST"`
	SmtpPort        int    `env:"EMAIL_SERVICE_SMTP_PORT"`
	NoreplyUsername string `env:"EMAIL_SERVICE_SMTP_NO-REPLY_USERNAME"`
	NoreplyPassword string `env:"EMAIL_SERVICE_SMTP_NO-REPLY_PASSWORD"`
	SupportUsername string `env:"EMAIL_SERVICE_SMTP_SUPPORT_USERNAME"`
	SupportPassword string `env:"EMAIL_SERVICE_SMTP_SUPPORT_PASSWORD"`
}