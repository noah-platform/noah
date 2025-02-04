package service

import (
	_ "embed"
	"html/template"
)

//go:embed assets/email_verification_template.tmpl
var emailVerificationTemplateString string
var emailVerificationTemplate = template.Must(template.New("email_verification").Parse(emailVerificationTemplateString))

//go:embed assets/email_password_reset_template.tmpl
var emailPasswordResetTemplateString string
var emailPasswordResetTemplate = template.Must(template.New("email_password_reset").Parse(emailPasswordResetTemplateString))

type EmailVerificationTemplateData struct {
	Name            string
	VerificationURL string
}

type EmailPasswordResetTemplateData struct {
	Name             string
	PasswordResetURL string
}
