package main

import (
	"github.com/rs/zerolog/log"
	"github.com/stephanvebrian/expense-manager/transaction-service/internal/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal().Err(err).Msg("error when initializing config")
	}

	log.Info().Msg("running application...")

	err = startApp(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("error when trying start the application")
	}
}
