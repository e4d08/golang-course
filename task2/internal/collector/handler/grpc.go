package handler

import (
	"context"
	"errors"
	"gh-anal/internal/collector/domain"
	"gh-anal/internal/collector/usecase"
	pb "gh-anal/internal/gen/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GrpcHandler struct {
	pb.UnimplementedCollectorServiceServer
	service usecase.CollectorService
}

func NewGrpcHandler(service usecase.CollectorService) *GrpcHandler {
	return &GrpcHandler{
		service: service,
	}
}

func (h *GrpcHandler) validateRequest(req *pb.GetRepositoryRequest) error {
	if req.Name == "" {
		return errors.New("parameter 'name' is required")
	}
	if req.Owner == "" {
		return errors.New("parameter 'owner' is required")
	}
	return nil
}

func (h *GrpcHandler) GetRepository(ctx context.Context, req *pb.GetRepositoryRequest) (*pb.GetRepositoryResponse, error) {
	if err := h.validateRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	repo, err := h.service.GetRepository(req.Name, req.Owner)
	switch {
	case err == nil:
		return &pb.GetRepositoryResponse{
			Name:        repo.Name,
			Description: repo.Description,
			StarsCount:  repo.Stars,
			ForksCount:  repo.Forks,
			CreatedAt:   timestamppb.New(repo.CreatedAt),
			License:     repo.License,
		}, nil
	case errors.Is(err, domain.ErrInternal):
		return nil, status.Error(codes.Internal, err.Error())
	case errors.Is(err, domain.ErrRepoNotFound):
		return nil, status.Error(codes.NotFound, err.Error())
	default:
		return nil, status.Error(codes.Unknown, "unknown error")
	}
}
