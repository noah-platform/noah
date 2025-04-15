package repository

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/noah-platform/noah/question-bank/server/core"
)

func (r *UserTestSessionRepository) GetManySessionsByUserID(ctx context.Context, userID string, isCompletedOnly bool) ([]core.UserTestSession, error) {
	l := log.Ctx(ctx)

	filter := bson.M{"userId": userID}
	opts := options.Find()
	if isCompletedOnly {
		filter = bson.M{"userId": userID, "completedAt": bson.M{"$ne": nil}}
		opts = opts.SetSort(bson.D{{"startedAt", -1}})
	}

	var sessions []core.UserTestSession
	cursor, err := r.userTestSession.Find(ctx, filter, opts)
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
