package grpcserver

import (
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
)

type ConfigDeps struct {
	Host    string
	Port    string
	Timeout time.Duration
}

type Server struct {
	host       string
	port       string
	timeout    time.Duration
	gRPCServer *grpc.Server
}

func NewServer(deps *ConfigDeps) *Server {
	s := grpc.NewServer()

	return &Server{
		host:       deps.Host,
		port:       deps.Port,
		timeout:    deps.Timeout,
		gRPCServer: s,
	}
}

func (s *Server) Run() error {
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.host, s.port))
	if err != nil {
		return err
	}

	if err := s.gRPCServer.Serve(l); err != nil {
		return err
	}

	return nil
}

// Stop - stops gRPC server
func (s *Server) Stop() {
	s.gRPCServer.GracefulStop()
}

func (s Server) Server() *grpc.Server {
	return s.gRPCServer
}
