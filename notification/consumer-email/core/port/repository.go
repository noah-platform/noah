package port

import (
	"context"

	"github.com/noah-platform/noah/notification/consumer-email/core"
)

type Mailer interface {
	Send(ctx context.Context, email core.OutgoingEmailMessage) error
}
