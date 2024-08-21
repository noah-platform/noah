package core

type OutgoingEmailMessage struct {
	From    string `json:"from" validate:"required,email"`
	To      string `json:"to" validate:"required,email"`
	Subject string `json:"subject" validate:"required"`
	Body    string `json:"body" validate:"required"`
}
