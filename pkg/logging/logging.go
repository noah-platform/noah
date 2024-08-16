package logging

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init() {
	if os.Getenv("APP_ENV") == "production" {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)

		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}
}
