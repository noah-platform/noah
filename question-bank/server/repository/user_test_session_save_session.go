package repository

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/noah-platform/noah/question-bank/server/core"
)

func (r *UserTestSessionRepository) SaveSession(ctx context.Context, session *core.UserTestSession) error {
	l := log.Ctx(ctx)

	filter := bson.M{"userId": session.UserID, "test.id": session.Test.ID}
	opts := options.Replace().SetUpsert(true)

	if _, err := r.userTestSession.ReplaceOne(ctx, filter, session, opts); err != nil {
		l.Error().Err(err).Msg("[UserTestSessionRepository.SaveSession] failed to save user test session")

		return errors.Wrap(err, "failed to save user test session")
	}

	l.Info().Msg("[UserTestSessionRepository.SaveSession] user test session saved")

	return nil
}
