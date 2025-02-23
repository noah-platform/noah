package repository

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/noah-platform/noah/question-bank/server/core"
)

func (r *QuestionBankRepository) GetTestById(ctx context.Context, id string) (*core.TestEntry, error) {
	l := log.Ctx(ctx)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		l.Error().Err(err).Msg("[Service.GetTestByID] failed to convert id to object id")

		return nil, errors.Wrap(err, "failed to convert id to object id")
	}

	var testDocument TestDocument
	if err := r.questionBank.FindOne(ctx, bson.M{"_id": objectID}).Decode(&testDocument); err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			l.Info().Msg("[QuestionBankRepository.GetTestById] test not found")

			return nil, core.ErrTestNotFound
		default:
			l.Error().Err(err).Msg("[QuestionBankRepository.GetTestById] failed to fetch test")

			return nil, errors.Wrap(err, "failed to fetch tests")
		}
	}

	l.Debug().Msg("[QuestionBankRepository.GetTestById] test loaded")

	return &core.TestEntry{
		ID:          testDocument.ID.Hex(),
		Title:       testDocument.Title,
		Duration:    testDocument.Duration,
		Instruction: testDocument.Instruction,
		Module:      testDocument.Module,
		Sections:    mapSections(testDocument.Sections),
	}, nil
}
