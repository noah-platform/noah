package repository

import (
	"github.com/wneessen/go-mail"
)

type Mailer struct {
	mailer *mail.Client
}

type MailerDependencies struct {
	Mailer *mail.Client
}

func NewMailer(deps MailerDependencies) *Mailer {
	return &Mailer{
		mailer: deps.Mailer,
	}
}
