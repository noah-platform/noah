package port

import (
	"context"

	"github.com/noah-platform/noah/auth/server-session/core"
)

type SessionRepository interface {
	GetSession(ctx context.Context, id string) (*core.Session, error)
	GetUserIDFromSession(ctx context.Context, id string) (string, error)
	CreateSession(ctx context.Context, session *core.Session) (*core.Session, error)
	DeleteSession(ctx context.Context, id string) error
}
