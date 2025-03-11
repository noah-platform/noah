package di

import (
	"github.com/noah-platform/noah/pkg/auth"
	"github.com/noah-platform/noah/pkg/validator"

	"github.com/noah-platform/noah/question-bank/server/core/service"
	"github.com/noah-platform/noah/question-bank/server/handler"
	"github.com/noah-platform/noah/question-bank/server/repository"
)

type ServerConfig = handler.Config

type ServiceConfig = service.Config

type WritingAssessmentRepoConfig = repository.WritingAssessmentRepoConfig

type Config struct {
	ServerConfig                ServerConfig
	ServiceConfig               ServiceConfig
	MongoConfig                 MongoConfig
	LambdaConfig                LambdaConfig
	WritingAssessmentRepoConfig WritingAssessmentRepoConfig
}

func New(cfg Config) *handler.Server {
	db := newMongoClient(cfg.MongoConfig)
	questionBankRepo := repository.NewQuestionBankRepository(repository.QuestionBankRepoDependencies{
		QuestionBank: db.Collection("question-bank"),
	})
	userTestSessionRepo := repository.NewUserTestSessionRepository(repository.UserTestSessionRepoDependencies{
		UserTestSession: db.Collection("user-test-sessions"),
	})

	lambdaClient := newLambdaClient(cfg.LambdaConfig)
	writingAssessmentRepo := repository.NewWritingAssessmentRepository(repository.WritingAssessmentRepoDependencies{
		LambdaClient: lambdaClient,
	}, cfg.WritingAssessmentRepoConfig)

	service := service.New(service.Dependencies{
		QuestionBankRepository:      questionBankRepo,
		UserTestSessionRepository:   userTestSessionRepo,
		WritingAssessmentRepository: writingAssessmentRepo,
	}, cfg.ServiceConfig)

	server := handler.New(handler.Dependencies{
		Service:        service,
		Validator:      validator.NewValidator(),
		AuthMiddleware: auth.NewMiddleware(),
	}, cfg.ServerConfig)

	return server
}
