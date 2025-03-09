package port

import (
	"context"

	"github.com/noah-platform/noah/question-bank/server/core"
)

type QuestionBankRepository interface {
	GetTestsByModule(ctx context.Context, module core.Module) ([]core.TestInfo, error)
	GetTestById(ctx context.Context, id string) (*core.TestEntry, error)
}

type UserTestSessionRepository interface {
	GetSession(ctx context.Context, userID, testID string) (*core.UserTestSession, error)
	GetManySessionsByUserID(ctx context.Context, userID string) ([]core.UserTestSession, error)
	SaveSession(ctx context.Context, session *core.UserTestSession) error
}
