package di

import (
	"github.com/noah-platform/noah/pkg/consumer"
	"github.com/noah-platform/noah/pkg/validator"

	"github.com/noah-platform/noah/notification/consumer-email/core/service"
	"github.com/noah-platform/noah/notification/consumer-email/handler"
)

type ConsumerConfig = consumer.Config

type HandlerConfig = handler.Config

type ServiceConfig = service.Config

type Config struct {
	ConsumerConfig ConsumerConfig
	HandlerConfig  HandlerConfig
	ServiceConfig  ServiceConfig
}

func New(cfg Config) *consumer.Consumer {
	service := service.New(service.Dependencies{}, cfg.ServiceConfig)

	handler := handler.New(handler.Dependencies{
		Service:   service,
		Validator: validator.NewValidator(),
	}, cfg.HandlerConfig)

	consumer := consumer.NewConsumer(consumer.Dependencies{
		Handler:   handler,
		Validator: validator.NewValidator(),
	}, cfg.ConsumerConfig)

	return consumer
}
