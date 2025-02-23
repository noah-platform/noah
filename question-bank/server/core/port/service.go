package port

import (
	"context"

	"github.com/noah-platform/noah/question-bank/server/core"
)

type Service interface {
	GetTestsByModule(ctx context.Context, module string) ([]core.TestInfo, error)
	GetTestById(ctx context.Context, id string) (*core.TestEntry, error)
}
