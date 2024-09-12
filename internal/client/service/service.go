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
	CancelOrder(ctx context.Context, orderID uint32) error
}

type serv struct {
	client pb.ClientServiceClient
}

func New(_client pb.ClientServiceClient) Service {
	return &serv{
		client: _client,
	}
}
