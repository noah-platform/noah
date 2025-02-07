package service

import (
	"bytes"
	"context"
	"crypto/rand"

	"github.com/lucsky/cuid"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"

	"github.com/noah-platform/noah/account/server/core"
	"github.com/noah-platform/noah/pkg/random"
)

func (s *Service) RegisterAccount(ctx context.Context, traceID, email, name, password string) (err error) {
	l := log.Ctx(ctx)

	userID, err := cuid.NewCrypto(rand.Reader)
	if err != nil {
		l.Error().Err(err).Msg("[Service.RegisterAccount] failed to generate userId")

		return errors.Wrap(err, "failed to generate userId")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		l.Error().Err(err).Msg("[Service.RegisterAccount] failed to hash password")

		return errors.Wrap(err, "failed to hash password")
	}
	passwordHash := string(hash)

	token, err := random.GenerateRandomString(64)
	if err != nil {
		l.Error().Err(err).Msg("[Service.RegisterAccount] failed to generate token")

		return errors.Wrap(err, "failed to generate token")
	}

	tx, err := s.accountRepo.BeginTransaction(ctx)
	if err != nil {
		l.Error().Err(err).Msg("[Service.RegisterAccount] failed to begin transaction")

		return errors.Wrap(err, "failed to begin transaction")
	}
	defer func() {
		if err != nil {
			if err := s.accountRepo.RollbackTransaction(ctx, tx); err == nil {
				l.Warn().Msg("[Service.RegisterAccount] transaction rolled back")
			} else {
				l.Error().Err(err).Msg("[Service.RegisterAccount] failed to rollback transaction")
			}
		}
	}()

	err = s.accountRepo.CreateAccount(ctx, &tx, core.Account{
		ID:                userID,
		Email:             email,
		Name:              name,
		Password:          &passwordHash,
		IsVerified:        false,
		VerificationToken: token,
	})
	if err != nil {
		switch {
		case errors.Is(err, core.ErrAccountAlreadyExists):
			l.Info().Msg("[Service.RegisterAccount] account already exists")

			return errors.Wrap(err, "account already exists")

		default:
			l.Error().Err(err).Msg("[Service.RegisterAccount] failed to create account")

			return errors.Wrap(err, "failed to create account")
		}
	}

	var body bytes.Buffer
	if err = emailVerificationTemplate.Execute(&body, EmailVerificationTemplateData{
		Name:            name,
		VerificationURL: "https://noah.example.com/verify/" + token,
	}); err != nil {
		l.Error().Err(err).Msg("[Service.RegisterAccount] failed to execute email verification template")

		return errors.Wrap(err, "failed to execute email verification template")
	}

	message := core.OutgoingEmailMessage{
		From:          s.config.EmailFrom,
		SenderName:    "Noah Platform",
		To:            email,
		RecipientName: name,
		Subject:       "Verify your email",
		Body:          body.String(),
	}
	if err := s.emailRepo.ProduceOutgoingEmail(ctx, traceID, message); err != nil {
		l.Error().Err(err).Msg("[Service.RegisterAccount] failed to produce email verification request")

		return errors.Wrap(err, "failed to produce email verification request")
	}

	if err := s.accountRepo.CommitTransaction(ctx, tx); err != nil {
		l.Error().Err(err).Msg("[Service.RegisterAccount] failed to commit transaction")

		return errors.Wrap(err, "failed to commit transaction")
	}

	l.Info().Msg("[Service.RegisterAccount] account created")

	return nil
}
