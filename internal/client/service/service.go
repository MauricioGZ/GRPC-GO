package service

import (
	"context"

	pb "github.com/MauricioGZ/GRPC-GO/internal/gen"
)

type Service interface {
	CreateOrder(ctx context.Context) (uint32, error)
	GetAllProducts(ctx context.Context) ([]*pb.Product, error)
}

type serv struct {
	client pb.OrdersServiceClient
}

func New(_client pb.OrdersServiceClient) Service {
	return &serv{
		client: _client,
	}
}
