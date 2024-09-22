package port

import (
	"context"
	"net/netip"

	"github.com/noah-platform/noah/auth/server-session/core"
)

type Service interface {
	GetSession(ctx context.Context, id string) (*core.Session, error)
	CreateSession(ctx context.Context, userID string, ipAddress netip.Addr, userAgent string) (*core.Session, error)
	DeleteSession(ctx context.Context, id string) error
	VerifySession(ctx context.Context, id string) (string, error)
}
