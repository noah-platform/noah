package port

import (
	"context"

	"github.com/noah-platform/noah/account/server/core"
)

type Service interface {
	GetAccount(ctx context.Context, id string) (*core.Account, error)
	GetAccountByEmail(ctx context.Context, email string) (*core.Account, error)
	GetAccountByGoogleAccountID(ctx context.Context, googleAccountID string) (*core.Account, error)
	RegisterAccount(ctx context.Context, traceID, email, name, password string) error
	VerifyEmail(ctx context.Context, token string) error
	RequestPasswordReset(ctx context.Context, traceID, email string) error
	ConfirmPasswordReset(ctx context.Context, token, password string) error
	RegisterAccountByGoogleAccount(ctx context.Context, email, name, googleAccountID string) (*core.Account, error)
}
