package service

import "github.com/noah-platform/noah/example/server/core/port"

type Service struct {
	config      Config
	exampleRepo port.ExampleRepository
}

type Config struct {
}

type Dependencies struct {
	ExampleRepository port.ExampleRepository
}

func New(deps Dependencies, cfg Config) *Service {
	return &Service{
		config:      cfg,
		exampleRepo: deps.ExampleRepository,
	}
}
