package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type ExampleMongoRepository struct {
	example *mongo.Collection
}

type ExampleMongoRepoDependencies struct {
	Example *mongo.Collection
}

func NewExampleMongoRepository(deps ExampleMongoRepoDependencies) *ExampleMongoRepository {
	return &ExampleMongoRepository{
		example: deps.Example,
	}
}
