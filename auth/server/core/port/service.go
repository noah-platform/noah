package port

import (
	"context"
)

type Service interface {
	Login(ctx context.Context, email, password string) (string, error)
	LoginWithGoogle(ctx context.Context, idToken string) (string, error)
	Logout(ctx context.Context, sessionID string) error
}
