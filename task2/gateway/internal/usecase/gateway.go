package usecase

import (
	"context"
	adapter "gateway/internal/adapter/grpc"
	"gateway/internal/domain"
)

type GatewayService struct {
	collector *adapter.GrpcAdapter
}

func NewGatewayService(collector *adapter.GrpcAdapter) *GatewayService {
	return &GatewayService{collector: collector}
}

func (s *GatewayService) GetRepository(ctx context.Context, owner string, name string) (domain.Repository, error) {
	repo, err := s.collector.GetRepository(ctx, owner, name)
	if err != nil {
		return domain.Repository{}, err
	}

	return repo, nil
}
