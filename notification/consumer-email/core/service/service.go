package service

type Service struct {
	config Config
}

type Config struct {
}

type Dependencies struct {
}

func New(deps Dependencies, cfg Config) *Service {
	return &Service{
		config: cfg,
	}
}
