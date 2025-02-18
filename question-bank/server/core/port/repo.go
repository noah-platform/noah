package port

import (
	"context"

	"github.com/noah-platform/noah/question-bank/server/core"
)

type QuestionBankRepository interface {
	GetTestsByModule(ctx context.Context, module string) ([]core.TestInfo, error)
	GetTestById(ctx context.Context, module string) (*core.TestEntry, error)
}
