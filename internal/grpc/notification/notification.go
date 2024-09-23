package notification

import (
	"context"

	ssov1 "github.com/Magic-Kot/store-protos/gen/go/sso"

	"google.golang.org/grpc"
)

type serverAPI struct {
	ssov1.UnimplementedNotificationServer
}

func Register(gRPC *grpc.Server) {
	ssov1.RegisterNotificationServer(gRPC, &serverAPI{})
}

func (s *serverAPI) TelegramNotification(ctx context.Context, in *ssov1.MessageRequest) (*ssov1.StatusResponse, error) {
	panic("implement me")
}
