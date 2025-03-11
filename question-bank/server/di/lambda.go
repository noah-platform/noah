package di

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

type LambdaConfig struct {
	Region          string
	AccessKeyID     string
	SecretAccessKey string
}

func newLambdaClient(cfg LambdaConfig) *lambda.Client {
	staticCredential := credentials.NewStaticCredentialsProvider(cfg.AccessKeyID, cfg.SecretAccessKey, "")

	config, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(cfg.Region), config.WithCredentialsProvider(staticCredential))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	lambdaClient := lambda.NewFromConfig(config)

	return lambdaClient
}
