package service

import (
	"context"

	"github.com/MauricioGZ/GRPC-GO/internal/client/api/dto"
	"github.com/MauricioGZ/GRPC-GO/internal/client/model"
	pb "github.com/MauricioGZ/GRPC-GO/internal/gen"
)

type Service interface {
	CreateOrder(ctx context.Context, orderItems []dto.OrderItem) (model.OrderResponse, error)
	GetAllProducts(ctx context.Context) ([]model.Product, error)
}

type serv struct {
	client pb.OrdersServiceClient
}

func New(_client pb.OrdersServiceClient) Service {
	return &serv{
		client: _client,
	}
}
