package service

import (
	"context"

	"github.com/noah-platform/noah/account/server/client"
	"github.com/noah-platform/noah/auth/server/core"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (s *Service) LoginWithGoogle(ctx context.Context, idToken string) (string, error) {
	l := log.Ctx(ctx)

	// TODO: replace with client id from config
	payload, err := s.googleIDTokenValidator.Validate(ctx, idToken, "1082322066020-0a2irp5i8b0shiq0njm8ksd5h4qvt4a8.apps.googleusercontent.com")
	if err != nil {
		l.Info().Err(err).Msg("[Service.LoginWithGoogle] failed to validate Google ID token")

		return "", core.ErrInvalidCredentials
	}

	*l = l.With().Str("googleAccountId", payload.Subject).Logger()

	var account *client.GetAccountByGoogleAccountIDResponse
	account, err = s.accountClient.FetchAccountByGoogleAccountID(payload.Subject)
	if err != nil {
		switch {
		case errors.Is(err, client.ErrAccountNotFound):
			account, err = s.accountClient.RegisterAccountByGoogleAccount(payload.Claims["email"].(string), payload.Claims["name"].(string), payload.Subject)
			if err != nil {
				l.Error().Err(err).Msg("[Service.LoginWithGoogle] failed to register account by google account id")

				return "", errors.Wrap(err, "failed to register account by google account id")
			}

			l.Info().Msg("[Service.LoginWithGoogle] new account registered")
		default:
			l.Error().Err(err).Msg("[Service.LoginWithGoogle] failed to fetch account by google account id")

			return "", errors.Wrap(err, "failed to fetch account by google account id")
		}
	}

	*l = l.With().Str("userId", account.ID).Logger()

	// TODO: call auth session server to create a new session

	l.Info().Str("sessionId", "TODO").Msg("[Service.LoginWithGoogle] login successfully")

	return "", nil
}
