package port

import (
	"context"
)

type Service interface {
	Login(ctx context.Context, email, password, ipAddress, userAgent string) (string, error)
	LoginWithGoogle(ctx context.Context, idToken string) (string, error)
	Logout(ctx context.Context, sessionID string) error
}
