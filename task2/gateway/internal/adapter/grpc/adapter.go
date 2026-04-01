package adapter

import (
	"context"
	"fmt"
	"gateway/internal/domain"
	pb "gateway/internal/gen/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type GrpcAdapter struct {
	conn   *grpc.ClientConn
	client pb.CollectorServiceClient
}

func NewGrpcAdapter(address string) (*GrpcAdapter, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("dial grpc: %w", err)
	}

	client := pb.NewCollectorServiceClient(conn)

	return &GrpcAdapter{conn: conn, client: client}, nil
}

func (a *GrpcAdapter) GetRepository(ctx context.Context, owner string, name string) (domain.Repository, error) {
	req := &pb.GetRepositoryRequest{
		Owner: owner,
		Name:  name,
	}

	resp, err := a.client.GetRepository(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			switch st.Code() {
			case codes.NotFound:
				return domain.Repository{}, domain.ErrRepositoryNotFound
			case codes.Internal:
				return domain.Repository{}, domain.ErrInternal
			case codes.InvalidArgument:
				return domain.Repository{}, fmt.Errorf("%w: %w", domain.ErrInvalidArgument, st.Err())
			default:
				return domain.Repository{}, err
			}
		} else {
			return domain.Repository{}, domain.ErrInternal
		}
	}

	return domain.Repository{
		Name:        resp.Name,
		Description: resp.Description,
		Stars:       resp.StarsCount,
		Forks:       resp.ForksCount,
		CreatedAt:   resp.CreatedAt.AsTime(),
		License:     resp.License,
	}, nil
}

func (a *GrpcAdapter) Close() {
	a.conn.Close()
}
