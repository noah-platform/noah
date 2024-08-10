package port

import (
	"context"

	"github.com/noah-platform/noah/notification/consumer-email/core"
)

type Service interface {
	SendEmail(ctx context.Context, email core.OutgoingEmailMessage) error
}
