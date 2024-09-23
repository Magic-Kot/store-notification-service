package grpcapp

import (
	"fmt"
	"net"

	"github.com/Magic-Kot/store-notification-service/internal/grpc/notification"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type App struct {
	log        *zerolog.Logger
	gRPCServer *grpc.Server
	port       int
}

// NewApp - creates new gRPC server app.
func NewApp(log *zerolog.Logger, port int) *App {
	gRPCServer := grpc.NewServer()

	notification.Register(gRPCServer)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}

}

func (a *App) Run() error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		a.log.Debug().Msgf("Error listening on port %d: %v", a.port, err)
		return err
	}

	a.log.Debug().Msgf("Starting gRPC server on addr: %s", l.Addr())

	if err := a.gRPCServer.Serve(l); err != nil {
		a.log.Debug().Msgf("Error starting gRPC server: %v", err)
		return err
	}

	return nil
}

// Stop - stops gRPC server
func (a *App) Stop() {
	a.log.Debug().Msg("Stopping gRPC server")

	a.gRPCServer.GracefulStop()
}
