package di

import (
	"github.com/rs/zerolog/log"
	"github.com/wneessen/go-mail"
)

type MailerConfig struct {
	SMTPHost string
	SMTPPort int
}

func newMailer(cfg MailerConfig) *mail.Client {
	// TODO: Implement SMTP authentication
	mailer, err := mail.NewClient(cfg.SMTPHost, mail.WithPort(cfg.SMTPPort), mail.WithSMTPAuth(mail.SMTPAuthNoAuth), mail.WithTLSPolicy(mail.NoTLS))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create mail client")
	}

	return mailer
}
