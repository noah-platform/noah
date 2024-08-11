package port

import (
	"context"

	"github.com/noah-platform/noah/account/server/core"
	"github.com/noah-platform/noah/pkg/transaction"
)

type AccountRepository interface {
	transaction.RepositoryWithTransaction
	GetAccount(ctx context.Context, id string) (*core.Account, error)
	CreateAccount(ctx context.Context, tx transaction.Tx, account core.Account) error
}

type EmailRepository interface {
	ProduceEmailVerificationRequest(ctx context.Context, to, name, url string) error
}
