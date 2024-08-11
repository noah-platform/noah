package port

import (
	"context"

	"github.com/noah-platform/noah/account/server/core"
)

type Service interface {
	GetAccount(ctx context.Context, id string) (*core.Account, error)
	RegisterAccount(ctx context.Context, traceID, email, name, password string) error
}
