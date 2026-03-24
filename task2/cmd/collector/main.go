package collector

import (
	"gh-anal/internal/collector/adapter/github"
	"gh-anal/internal/collector/handler"
	"gh-anal/internal/collector/usecase"
	pb "gh-anal/internal/gen/proto"
	"net"
	"time"

	"google.golang.org/grpc"
)

func main() {
	githubAdapter := github.NewHttpAdapter(5 * time.Second)
	service := usecase.NewCollectorService(githubAdapter)
	grpcHandler := handler.NewGrpcHandler(service)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic("failed to listen: " + err.Error())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCollectorServiceServer(grpcServer, grpcHandler)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			panic("failed to serve: " + err.Error())
		}
	}()
}
