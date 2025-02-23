package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type QuestionBankRepository struct {
	questionBank *mongo.Collection
}

type QuestionBankRepoDependencies struct {
	QuestionBank *mongo.Collection
}

func NewQuestionBankRepository(deps QuestionBankRepoDependencies) *QuestionBankRepository {
	return &QuestionBankRepository{
		questionBank: deps.QuestionBank,
	}
}
