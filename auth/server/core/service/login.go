package service

import (
	"context"

	"github.com/noah-platform/noah/account/server/client"
	"github.com/noah-platform/noah/auth/server/core"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Login(ctx context.Context, email, password string) (string, error) {
	l := log.Ctx(ctx)
	*l = l.With().Str("email", email).Logger()

	account, err := s.accountClient.FetchAccountByEmail(email)
	if err != nil {
		switch {
		case errors.Is(err, client.ErrAccountNotFound):
			l.Info().Err(err).Msg("[Service.Login] account not found")

			return "", core.ErrInvalidCredentials
		default:
			l.Error().Err(err).Msg("[Service.Login] failed to fetch account by email")

			return "", errors.Wrap(err, "failed to fetch account by email")
		}
	}

	*l = l.With().Str("userId", account.ID).Logger()

	if account.GoogleAccountID != nil {
		l.Info().Msg("[Service.Login] account is a Google account, login by email is denied")

		return "", core.ErrInvalidCredentials
	}

	if account.Password == nil {
		l.Error().Msg("[Service.Login] account has no password")

		return "", errors.New("account has no password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(*account.Password), []byte(password))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			l.Info().Err(err).Msg("[Service.Login] invalid password")

			return "", core.ErrInvalidCredentials
		default:
			l.Error().Err(err).Msg("[Service.Login] failed to compare password")

			return "", errors.Wrap(err, "failed to compare password")
		}
	}

	if !account.IsVerified {
		l.Info().Msg("[Service.Login] account is not verified")

		return "", core.ErrAccountNotVerified
	}

	// TODO: call auth session server to create a new session

	return "", nil
}
