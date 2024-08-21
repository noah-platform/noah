package repository

import (
	"github.com/noah-platform/noah/pkg/producer"
)

type EmailRepository struct {
	producer  *producer.Producer
	topic     string
	emailFrom string
}

type EmailRepoDependencies struct {
	Producer *producer.Producer
}

type EmailRepoConfig struct {
	KafkaTopic string
	EmailFrom  string
}

func NewEmailRepository(deps EmailRepoDependencies, cfg EmailRepoConfig) *EmailRepository {
	return &EmailRepository{
		producer:  deps.Producer,
		topic:     cfg.KafkaTopic,
		emailFrom: cfg.EmailFrom,
	}
}
