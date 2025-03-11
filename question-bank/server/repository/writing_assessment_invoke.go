package repository

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/question-bank/server/core"
)

func (r *WritingAssessmentRepository) Invoke(ctx context.Context, input *core.WritingAssessmentInput) (*core.WritingAssessmentOutput, error) {
	l := log.Ctx(ctx)
	*l = l.With().Interface("input", input).Logger()

	body, err := json.Marshal(input)
	if err != nil {
		l.Err(err).Msg("[WritingAssessmentRepository.AssessTaskOne] failed to marshal body")

		return nil, errors.Wrap(err, "failed to marshal body")
	}

	payload, err := json.Marshal(LambdaRequest{
		Body: string(body),
	})
	if err != nil {
		l.Err(err).Msg("[WritingAssessmentRepository.AssessTaskOne] failed to marshal payload")

		return nil, errors.Wrap(err, "failed to marshal payload")
	}

	invokeOutput, err := r.lambdaClient.Invoke(ctx, &lambda.InvokeInput{
		FunctionName: aws.String(r.functionName),
		LogType:      types.LogTypeNone,
		Payload:      payload,
	})
	if err != nil {
		l.Err(err).Msg("[WritingAssessmentRepository.AssessTaskOne] failed to invoke lambda function")

		return nil, errors.Wrap(err, "failed to invoke lambda function")
	}

	var response LambdaResponse
	if err := json.Unmarshal(invokeOutput.Payload, &response); err != nil {
		l.Err(err).Msg("[WritingAssessmentRepository.AssessTaskOne] failed to unmarshal response")

		return nil, errors.Wrap(err, "failed to unmarshal response")
	}

	var output core.WritingAssessmentOutput
	if err := json.Unmarshal([]byte(response.Body), &output); err != nil {
		l.Err(err).Msg("[WritingAssessmentRepository.AssessTaskOne] failed to unmarshal output")

		return nil, errors.Wrap(err, "failed to unmarshal output")
	}

	l.Info().Interface("output", invokeOutput).Msg("[WritingAssessmentRepository.AssessTaskOne] invoked lambda function")

	return &output, nil
}
