package main

import (
	"github.com/caarlos0/env/v11"
	"github.com/rs/zerolog/log"

	"github.com/noah-platform/noah/pkg/logging"
	"github.com/noah-platform/noah/question-bank/server/di"
)

type Config struct {
	Environment string `env:"APP_ENV,required"`

	Port      string `env:"PORT,required"`
	JWTSecret string `env:"JWT_SECRET,required"`
	MongoUri  string `env:"MONGO_URI,required"`

	LambdaRegion          string `env:"LAMBDA_REGION,required"`
	LambdaAccessKeyID     string `env:"LAMBDA_ACCESS_KEY_ID,required"`
	LambdaSecretAccessKey string `env:"LAMBDA_SECRET_ACCESS_KEY,required"`

	WritingAssessmentLambdaFunctionName string `env:"WRITING_ASSESSMENT_LAMBDA_FUNCTION_NAME,required"`
}

func init() {
	logging.Init()
}

func main() {
	cfg, err := env.ParseAs[Config]()
	if err != nil {
		log.Panic().Err(err).Msg("failed to load config")
	}

	log.Info().Str("environment", cfg.Environment).Msg("server is starting")

	server := di.New(di.Config{
		ServerConfig: di.ServerConfig{
			Port:      cfg.Port,
			JWTSecret: cfg.JWTSecret,
		},
		ServiceConfig: di.ServiceConfig{},
		MongoConfig: di.MongoConfig{
			DatabaseUri: cfg.MongoUri,
		},
		LambdaConfig: di.LambdaConfig{
			Region:          cfg.LambdaRegion,
			AccessKeyID:     cfg.LambdaAccessKeyID,
			SecretAccessKey: cfg.LambdaSecretAccessKey,
		},
		WritingAssessmentRepoConfig: di.WritingAssessmentRepoConfig{
			FunctionName: cfg.WritingAssessmentLambdaFunctionName,
		},
	})

	server.Start()
}
