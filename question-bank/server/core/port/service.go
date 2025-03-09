package port

import (
	"context"

	"github.com/noah-platform/noah/question-bank/server/core"
)

type Service interface {
	GetTestsByModule(ctx context.Context, userID string, module core.Module) ([]core.UserTestInfo, error)
	GetTestCover(ctx context.Context, id string) (*core.TestCover, error)
	BeginTestSession(ctx context.Context, userID, testID string) (*core.UserTestSession, error)
	SaveTestSession(ctx context.Context, userID, testID string, answers map[string]string) error
	EndTestSession(ctx context.Context, userID, testID string, answers map[string]string) error
}
