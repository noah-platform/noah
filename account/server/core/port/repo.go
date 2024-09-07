package port

import (
	"context"

	"github.com/noah-platform/noah/account/server/core"
	"github.com/noah-platform/noah/pkg/transaction"
)

type AccountRepository interface {
	transaction.RepositoryWithTransaction
	GetAccount(ctx context.Context, id string) (*core.Account, error)
	GetAccountByEmail(ctx context.Context, email string) (*core.Account, error)
	CreateAccount(ctx context.Context, tx transaction.Tx, account core.Account) error
	CreatePasswordResetToken(ctx context.Context, token core.PasswordResetToken) error
}

type EmailRepository interface {
	ProduceOutgoingEmail(ctx context.Context, traceID string, message core.OutgoingEmailMessage) error
}
