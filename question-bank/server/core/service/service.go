package service

import "github.com/noah-platform/noah/question-bank/server/core/port"

type Service struct {
	config                Config
	questionBankRepo      port.QuestionBankRepository
	userTestSessionRepo   port.UserTestSessionRepository
	writingAssessmentRepo port.WritingAssessmentRepository
}

type Config struct {
}

type Dependencies struct {
	QuestionBankRepository      port.QuestionBankRepository
	UserTestSessionRepository   port.UserTestSessionRepository
	WritingAssessmentRepository port.WritingAssessmentRepository
}

func New(deps Dependencies, cfg Config) *Service {
	return &Service{
		config:                cfg,
		questionBankRepo:      deps.QuestionBankRepository,
		userTestSessionRepo:   deps.UserTestSessionRepository,
		writingAssessmentRepo: deps.WritingAssessmentRepository,
	}
}
