package service

import (
	_ "embed"
	"html/template"
)

//go:embed assets/email_verification_template.tmpl
var emailVerificationTemplateString string
var emailVerificationTemplate = template.Must(template.New("email_verification").Parse(emailVerificationTemplateString))

type EmailVerificationTemplateData struct {
	Name            string
	VerificationURL string
}
