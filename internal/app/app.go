package app

import (
	grpcapp "github.com/Magic-Kot/store-notification-service/internal/app/grpc"

	"github.com/rs/zerolog"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func NewAppServ(log *zerolog.Logger, grpcPort int) *App {
	grpcApp := grpcapp.NewApp(log, grpcPort)
	return &App{
		GRPCSrv: grpcApp,
	}
}
