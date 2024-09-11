package service

import (
	"context"

	"github.com/MauricioGZ/GRPC-GO/internal/client/api/dto"
	"github.com/MauricioGZ/GRPC-GO/internal/client/model"
	pb "github.com/MauricioGZ/GRPC-GO/internal/gen"
)

func (s *serv) CreateOrder(ctx context.Context, orderItems []dto.OrderItem) (model.OrderResponse, error) {
	var orderItemsRequest []*pb.OrderItem

	for _, orderItem := range orderItems {
		orderItemsRequest = append(orderItemsRequest,
			&pb.OrderItem{
				ProductID: orderItem.ProductID,
				Quantity:  orderItem.Quantity,
			})
	}

	orderID, err := s.client.CreateOrder(ctx, &pb.CreateOrderRequest{
		CustomerID: 1,
		OrderItems: orderItemsRequest,
	})
	return model.OrderResponse{OrderID: orderID.GetOrderID()}, err
}
