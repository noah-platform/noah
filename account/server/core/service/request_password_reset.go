package service

import (
	"bytes"
	"context"
	"net/url"
	"time"

	"github.com/noah-platform/noah/account/server/core"
	"github.com/noah-platform/noah/pkg/random"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (s *Service) RequestPasswordReset(ctx context.Context, traceID, email string) (err error) {
	l := log.Ctx(ctx)
	*l = l.With().Str("traceID", traceID).Str("email", email).Logger()
	ctx = l.WithContext(ctx)

	account, err := s.accountRepo.GetAccountByEmail(ctx, email)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrAccountNotFound):
			l.Info().Err(err).Msg("[Service.RequestPasswordReset] account not found")

			return nil
		default:
			l.Error().Err(err).Msgf("[Service.RequestPasswordReset] failed to get account")

			return errors.Wrap(err, "failed to get account")
		}
	}

	token, err := random.GenerateRandomString(64)
	if err != nil {
		l.Error().Err(err).Msg("[Service.RequestPasswordReset] failed to generate token")

		return errors.Wrap(err, "failed to generate token")
	}

	err = s.accountRepo.CreatePasswordResetToken(ctx, core.PasswordResetToken{
		Token:     token,
		UserID:    account.ID,
		ExpiresAt: time.Now().Add(time.Hour),
	})
	if err != nil {
		l.Error().Err(err).Msg("[Service.RequestPasswordReset] failed to create password reset token")

		return errors.Wrap(err, "failed to create password reset token")
	}

	passwordResetURL, err := url.JoinPath(s.config.FrontendBaseUrl, "forget-password", token)
	if err != nil {
		l.Error().Err(err).Msg("[Service.RequestPasswordReset] failed to join password reset URL")

		return errors.Wrap(err, "failed to join password reset URL")
	}

	var body bytes.Buffer
	if err = emailPasswordResetTemplate.Execute(&body, EmailPasswordResetTemplateData{
		Name:             account.Name,
		PasswordResetURL: passwordResetURL,
	}); err != nil {
		l.Error().Err(err).Msg("[Service.RequestPasswordReset] failed to execute email password reset template")

		return errors.Wrap(err, "failed to execute email password reset template")
	}

	message := core.OutgoingEmailMessage{
		From:          s.config.EmailFrom,
		SenderName:    "Noah Platform",
		To:            email,
		RecipientName: account.Name,
		Subject:       "Reset your password",
		Body:          body.String(),
	}
	if err := s.emailRepo.ProduceOutgoingEmail(ctx, traceID, message); err != nil {
		l.Error().Err(err).Msg("[Service.RequestPasswordReset] failed to produce email password reset request")

		return errors.Wrap(err, "failed to produce email password reset request")
	}

	l.Info().Msg("[Service.RequestPasswordReset] password reset token created")

	return nil
}
