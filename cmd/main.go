package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Magic-Kot/store-notification-service/internal/config"
	"github.com/Magic-Kot/store-notification-service/internal/grpc/notification"
	"github.com/Magic-Kot/store-notification-service/internal/services/email"
	"github.com/Magic-Kot/store-notification-service/pkg/grpcserver"
	"github.com/Magic-Kot/store-notification-service/pkg/logging"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/rs/zerolog/log"
)

func main() {
	// read config
	var cfg config.Config

	err := cleanenv.ReadConfig("internal/config/config.yml", &cfg) // Local: internal/config/config.yml, Docker: config.yml
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
	serv := grpcserver.ConfigDeps{
		Host:    cfg.ServerDeps.Host,
		Port:    cfg.ServerDeps.Port,
		Timeout: cfg.ServerDeps.Timeout,
	}

	gRPCServer := grpcserver.NewServer(&serv)

	// email
	emailService := email.NewEmailService(&cfg.MailService)
	notification.Register(gRPCServer.Server(), emailService, logger)

	// start server
	go func() {
		err = gRPCServer.Run()
		if err != nil {
			logger.Fatal().Err(err).Msg("error starting gRPC server")
		}
	}()
	logger.Info().Msgf("starting gRPC server on addr: %s:%s", cfg.ServerDeps.Host, cfg.ServerDeps.Port)

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	gRPCServer.Stop()
	logger.Info().Msg("gracefully stopped")
}
