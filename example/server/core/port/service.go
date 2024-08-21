package port

import (
	"context"

	"github.com/noah-platform/noah/example/server/core"
)

type Service interface {
	GetExample(ctx context.Context, id int) (*core.Example, error)
}
