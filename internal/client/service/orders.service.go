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

	createOrderResponse, err := s.client.CreateOrder(ctx, &pb.CreateOrderRequest{
		CustomerID: 1,
		OrderItems: orderItemsRequest,
	})
	return model.OrderResponse{OrderID: createOrderResponse.GetOrderID()}, err
}

func (s *serv) CancelOrder(ctx context.Context, orderID uint32) error {
	_, err := s.client.CancelOrder(ctx, &pb.CancelOrderRequest{OrderID: orderID})
	return err
}
