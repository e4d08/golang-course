package grpc

import (
	"collector/internal/domain"
	pb "collector/internal/gen/proto"
	"collector/internal/usecase"
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Handler struct {
	pb.UnimplementedCollectorServiceServer
	service *usecase.CollectorService
}

func NewHandler(service *usecase.CollectorService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetRepository(ctx context.Context, req *pb.GetRepositoryRequest) (*pb.GetRepositoryResponse, error) {
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}
	if req.Owner == "" {
		return nil, status.Error(codes.InvalidArgument, "owner is required")
	}

	repo, err := h.service.GetRepository(ctx, req.Owner, req.Name)
	if err != nil {
		return nil, mapError(err)
	}

	return &pb.GetRepositoryResponse{
		Name:        repo.Name,
		Description: repo.Description,
		StarsCount:  repo.Stars,
		ForksCount:  repo.Forks,
		CreatedAt:   timestamppb.New(repo.CreatedAt),
		License:     repo.License,
	}, nil
}

func mapError(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, domain.ErrRepositoryNotFound) {
		return status.Error(codes.NotFound, "repository not found")
	}
	if errors.Is(err, domain.ErrInternal) {
		return status.Error(codes.Internal, "internal server error")
	}
	return status.Error(codes.Unknown, "unknown error, please try again later maybe")
}
