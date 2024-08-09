package di

import (
	"context"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	DatabaseUrl  string
	DatabaseName string
}

func newMongoClient(cfg MongoConfig) *mongo.Database {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.DatabaseUrl))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to mongo")
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatal().Err(err).Msg("failed to ping mongo")
	}

	db := client.Database(cfg.DatabaseName)

	return db
}
