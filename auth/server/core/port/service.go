package port

import (
	"context"

	"github.com/noah-platform/noah/auth/server/core"
)

type Service interface {
	Login(ctx context.Context, email, password, ipAddress, userAgent string) (string, error)
	LoginWithGoogle(ctx context.Context, idToken, ipAddress, userAgent string) (string, error)
	Logout(ctx context.Context, sessionID string) error
	GetMe(ctx context.Context, userID string) (*core.Me, error)
}
