package port

import (
	"context"

	"github.com/noah-platform/noah/question-bank/server/core"
)

type Service interface {
	GetTestsByModule(ctx context.Context, userID, module string) ([]core.UserTestInfo, error)
	GetTestById(ctx context.Context, id string) (*core.TestEntry, error)
	BeginTestSession(ctx context.Context, userID, testID string) (*core.UserTestSession, error)
	SaveTestSession(ctx context.Context, userID, testID string, answers map[string]string) error
	EndTestSession(ctx context.Context, userID, testID string, answers map[string]string) error
}
