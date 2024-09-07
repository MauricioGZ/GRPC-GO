package service

import (
	"context"

	pb "github.com/MauricioGZ/GRPC-GO/internal/gen"
)

func (s *serv) CreateOrder(ctx context.Context) (uint32, error) {
	orderID, err := s.client.CreateOrder(ctx, &pb.CreateOrderRequest{
		CustomerID: 1,
		OrderItems: []*pb.OrderItem{
			{
				ProductID: 1,
				Quantity:  2,
			},
			{
				ProductID: 2,
				Quantity:  2,
			},
			{
				ProductID: 3,
				Quantity:  2,
			},
		},
	})
	return orderID.GetOrderID(), err
}
