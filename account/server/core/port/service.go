package port

import (
	"context"

	"github.com/noah-platform/noah/account/server/core"
)

type Service interface {
	GetAccount(ctx context.Context, id string) (*core.Account, error)
}
