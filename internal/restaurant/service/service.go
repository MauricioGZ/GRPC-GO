package service

import (
	"context"

	pb "github.com/MauricioGZ/GRPC-GO/internal/gen"
)

type Service interface {
	GetPendingOrders(ctx context.Context) ([]*pb.GetPendingOrdersResponse, error)
}

type serv struct {
	client pb.RestaurantServiceClient
}

func New(_client pb.RestaurantServiceClient) Service {
	return &serv{
		client: _client,
	}
}
