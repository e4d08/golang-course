package grpc

import (
	pb "collector/internal/gen/proto"
	"collector/internal/usecase"
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	server *grpc.Server
	port   int
}

func NewServer(port int, service *usecase.CollectorService) *Server {
	server := grpc.NewServer()
	handler := NewHandler(service)
	pb.RegisterCollectorServiceServer(server, handler)

	return &Server{server: server, port: port}
}

func (s *Server) Run(ctx context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}

	go func() {
		s.server.Serve(lis)
	}()

	<-ctx.Done()

	s.server.GracefulStop()
	return nil
}
