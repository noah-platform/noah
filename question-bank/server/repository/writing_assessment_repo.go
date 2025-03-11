package repository

import (
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

type WritingAssessmentRepository struct {
	lambdaClient *lambda.Client
	functionName string
}

type WritingAssessmentRepoDependencies struct {
	LambdaClient *lambda.Client
}

type WritingAssessmentRepoConfig struct {
	FunctionName string
}

func NewWritingAssessmentRepository(deps WritingAssessmentRepoDependencies, cfg WritingAssessmentRepoConfig) *WritingAssessmentRepository {
	return &WritingAssessmentRepository{
		lambdaClient: deps.LambdaClient,
		functionName: cfg.FunctionName,
	}
}
