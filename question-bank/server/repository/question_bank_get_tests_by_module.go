package repository

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/noah-platform/noah/question-bank/server/core"
)

func (r *QuestionBankRepository) GetTestsByModule(ctx context.Context, module string) ([]core.TestInfo, error) {
	l := log.Ctx(ctx)

	var testDocuments []TestDocument
	cursor, err := r.questionBank.Find(ctx, bson.M{"module": module})
	if err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			l.Info().Msg("[QuestionBankRepository.GetTestsByModule] tests not found")

			return nil, core.ErrTestNotFound
		default:
			l.Error().Err(err).Msg("[QuestionBankRepository.GetTestsByModule] failed to fetch tests")

			return nil, errors.Wrap(err, "failed to fetch tests")
		}
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &testDocuments); err != nil {
		l.Error().Err(err).Msg("[QuestionBankRepository.GetTestsByModule] failed to decode tests")

		return nil, errors.Wrap(err, "failed to decode tests")
	}

	l.Debug().Msg("[QuestionBankRepository.GetTestsByModule] tests loaded")

	testInfoList := make([]core.TestInfo, 0, len(testDocuments))
	for _, doc := range testDocuments {
		testInfoList = append(testInfoList, core.TestInfo{
			ID:     doc.ID.Hex(),
			Module: doc.Module,
		})
	}

	return testInfoList, nil
}
