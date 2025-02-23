package di

import (
	"github.com/noah-platform/noah/pkg/validator"

	"github.com/noah-platform/noah/question-bank/server/core/service"
	"github.com/noah-platform/noah/question-bank/server/handler"
	"github.com/noah-platform/noah/question-bank/server/repository"
)

type ServerConfig = handler.Config

type ServiceConfig = service.Config

type Config struct {
	ServerConfig  ServerConfig
	ServiceConfig ServiceConfig
	MongoConfig   MongoConfig
}

func New(cfg Config) *handler.Server {

	db := newMongoClient(cfg.MongoConfig)
	questionBankRepo := repository.NewQuestionBankRepository(repository.QuestionBankRepoDependencies{
		QuestionBank: db.Collection("question-bank"),
	})

	service := service.New(service.Dependencies{
		QuestionBankRepository: questionBankRepo,
	}, cfg.ServiceConfig)

	server := handler.New(handler.Dependencies{
		Service:   service,
		Validator: validator.NewValidator(),
	}, cfg.ServerConfig)

	return server
}
