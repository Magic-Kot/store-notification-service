package main

import (
	"context"

	"github.com/Magic-Kot/store-notification-service/internal/app"
	"github.com/Magic-Kot/store-notification-service/internal/config"
	"github.com/Magic-Kot/store-notification-service/pkg/logging"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/rs/zerolog/log"
)

func main() {
	// read config
	var cfg config.Config

	err := cleanenv.ReadConfig("internal/config/config.yml", &cfg) // Local: internal/config/config.yml Docker: config.yml
	if err != nil {
		log.Fatal().Err(err).Msg("error initializing config")
	}

	// create logger
	logCfg := logging.LoggerDeps{
		LogLevel: cfg.LoggerDeps.LogLevel,
	}

	logger, err := logging.NewLogger(&logCfg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to init logger")
	}

	logger.Info().Msg("init logger")

	ctx := context.Background()
	ctx = logger.WithContext(ctx)

	logger.Debug().Msgf("config: %+v", cfg)

	// create server
	application := app.NewAppServ(logger, cfg.GRPC.Port)

	application.GRPCSrv.Run()
}
