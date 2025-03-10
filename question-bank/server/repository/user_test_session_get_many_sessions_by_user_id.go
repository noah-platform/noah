package repository

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/noah-platform/noah/question-bank/server/core"
)

func (r *UserTestSessionRepository) GetManySessionsByUserID(ctx context.Context, userID string) ([]core.UserTestSession, error) {
	l := log.Ctx(ctx)

	var sessions []core.UserTestSession
	cursor, err := r.userTestSession.Find(ctx, bson.M{"userId": userID})
	if err != nil {
		l.Error().Err(err).Msg("[UserTestSessionRepository.GetManySessionsByUserID] failed to get user test session")

		return nil, errors.Wrap(err, "failed to get user test session")
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &sessions); err != nil {
		l.Error().Err(err).Msg("[UserTestSessionRepository.GetManySessionsByUserID] failed to decode user test sessions")

		return nil, errors.Wrap(err, "failed to decode user test sessions")
	}

	l.Debug().Msg("[UserTestSessionRepository.GetManySessionsByUserID] test sessions loaded")

	return sessions, nil
}
