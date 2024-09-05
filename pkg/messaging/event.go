package messaging

type Event string

const (
	EventOutgoingEmail Event = "outgoing-email"
)

type OutgoingEmailMessage struct {
	From          string `json:"from" validate:"required,email"`
	SenderName    string `json:"senderName" validate:"required"`
	To            string `json:"to" validate:"required,email"`
	RecipientName string `json:"recipientName" validate:"required"`
	Subject       string `json:"subject" validate:"required"`
	Body          string `json:"body" validate:"required"`
}
