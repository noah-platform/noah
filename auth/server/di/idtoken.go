package di

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/api/idtoken"
)

func newGoogleIDTokenValidator() *idtoken.Validator {
	validator, err := idtoken.NewValidator(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create google id token validator")
	}

	return validator
}
