package di

import (
	"context"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

type MongoConfig struct {
	DatabaseUri string
}

func newMongoClient(cfg MongoConfig) *mongo.Database {
	connString, err := connstring.ParseAndValidate(cfg.DatabaseUri)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse mongo uri")
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connString.Original))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to mongo")
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatal().Err(err).Msg("failed to ping mongo")
	}

	db := client.Database(connString.Database)

	return db
}
