package repository

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/noah-platform/noah/example/server/core"
)

func (r *ExampleMongoRepository) GetExample(ctx context.Context, id int) (*core.Example, error) {
	l := log.Ctx(ctx)

	var example ExampleDocument
	if err := r.example.FindOne(ctx, bson.M{"exampleId": id}).Decode(&example); err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			l.Info().Msg("[ExampleMongoRepository.GetExample] example not found")

			return nil, core.ErrExampleNotFound
		default:
			l.Error().Err(err).Msg("[ExampleMongoRepository.GetExample] failed to fetch example")

			return nil, errors.Wrap(err, "failed to fetch example")
		}
	}

	l.Debug().Msg("[ExampleMongoRepository.GetExample] example loaded")

	return &core.Example{
		ID:    example.ID,
		Title: example.Title,
	}, nil
}
