package port

import (
	"context"

	"github.com/noah-platform/noah/account/server/core"
)

type Service interface {
	GetAccount(ctx context.Context, id string) (*core.Account, error)
	GetAccountByEmail(ctx context.Context, email string) (*core.Account, error)
	RegisterAccount(ctx context.Context, traceID, email, name, password string) error
	RequestPasswordReset(ctx context.Context, traceID, email string) error
	ConfirmPasswordReset(ctx context.Context, token, password string) error
}
