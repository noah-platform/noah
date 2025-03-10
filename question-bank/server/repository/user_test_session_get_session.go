package repository

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/noah-platform/noah/question-bank/server/core"
)

func (r *UserTestSessionRepository) GetSession(ctx context.Context, userID, testID string) (*core.UserTestSession, error) {
	l := log.Ctx(ctx)

	var session core.UserTestSession
	if err := r.userTestSession.FindOne(ctx, bson.M{"userId": userID, "test.id": testID}).Decode(&session); err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			l.Debug().Msg("[UserTestSessionRepository.GetSession] test not found")

			return nil, core.ErrUserTestSessionNotFound
		default:
			l.Error().Err(err).Msg("[UserTestSessionRepository.GetSession] failed to get user test session")

			return nil, errors.Wrap(err, "failed to get user test session")
		}
	}

	l.Debug().Msg("[UserTestSessionRepository.GetSession] user test session loaded")

	return &session, nil
}
