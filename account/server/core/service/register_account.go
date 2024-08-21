package service

import (
	"context"
	"crypto/rand"

	"github.com/lucsky/cuid"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"

	"github.com/noah-platform/noah/account/server/core"
)

func (s *Service) RegisterAccount(ctx context.Context, email, name, password string) error {
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

	err = s.accountRepo.CreateAccount(ctx, core.Account{
		ID:         userID,
		Email:      email,
		Name:       name,
		Password:   &passwordHash,
		IsVerified: false,
	})
	if err != nil {
		switch {
		case errors.Is(err, core.ErrAccountAlreadyExists):
			l.Info().Msg("[Service.RegisterAccount] account already exists")

			return errors.Wrap(err, "account already exists")

		default:
			l.Error().Err(err).Msgf("[Service.RegisterAccount] failed to create account")

			return errors.Wrap(err, "failed to create account")
		}
	}

	l.Info().Msg("[Service.RegisterAccount] account created")

	return nil
}
