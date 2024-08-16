package service

import _ "embed"

//go:embed assets/email_verification_template.tmpl
var emailVerificationTemplate string

type EmailVerificationTemplateData struct {
	Name            string
	VerificationURL string
}
