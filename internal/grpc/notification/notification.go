package notification

import (
	"context"

	ssov1 "github.com/Magic-Kot/store-protos/gen/go/sso"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SendEmailService interface {
	SendEmail(ctx context.Context, email string, subject string, body string) error
}

type serverAPI struct {
	ssov1.UnimplementedNotificationServer
	notification SendEmailService
	logger       *zerolog.Logger
}

func Register(gRPC *grpc.Server, notification SendEmailService, logger *zerolog.Logger) {
	ssov1.RegisterNotificationServer(gRPC, &serverAPI{notification: notification, logger: logger})
}

func (s *serverAPI) TelegramNotification(ctx context.Context, in *ssov1.MessageRequest) (*ssov1.StatusResponse, error) {
	ctx = s.logger.WithContext(ctx)
	s.logger.Debug().Msg("starting the handler 'TelegramNotification'")

	// validation
	if in.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	if in.Subject == "" {
		return nil, status.Error(codes.InvalidArgument, "message is required")
	}

	if in.Body == "" {
		return nil, status.Error(codes.InvalidArgument, "message is required")
	}

	err := s.notification.SendEmail(ctx, in.Email, in.Subject, in.Body)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	//формирование ответа
	return &ssov1.StatusResponse{
		Status: "200",
	}, nil
}
