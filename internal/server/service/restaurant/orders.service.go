package restaurant_service

import (
	"context"

	pb "github.com/MauricioGZ/GRPC-GO/internal/gen"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *service) GetPendingOrders(req *pb.GetPendingOrdersRequest, stream pb.RestaurantService_GetPendingOrdersServer) error {
	orders, err := s.repo.GetPendingOrders(context.Background())
	if err != nil {
		return err
	}

	for _, order := range orders {
		orderItems, err := s.wrapOrderItems(context.Background(), order.ID)
		if err != nil {
			return err
		}
		res := &pb.GetPendingOrdersResponse{
			OrderID:                order.ID,
			CustomerID:             order.CustomerID,
			OrderDate:              timestamppb.New(order.OrderDate),
			OrderItemByProductName: orderItems,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
	return nil
}

func (s *service) SetOrderToReady(ctx context.Context, req *pb.SetOrderToReadyRequest) (*pb.SetOrderToReadyResponse, error) {
	//check if the order does exist
	_, err := s.repo.GetOrderByID(ctx, req.OrderID)
	if err != nil {
		return nil, err
	}

	err = s.repo.UpdateOrderStatus(ctx, "ready", req.OrderID)
	if err != nil {
		return nil, err
	}
	return &pb.SetOrderToReadyResponse{}, nil
}

func (s *service) wrapOrderItems(ctx context.Context, orderID uint32) ([]*pb.OrderItemByProductName, error) {
	var orderItems []*pb.OrderItemByProductName
	ooii, err := s.repo.GetOrderItemsByOrderID(ctx, orderID)
	if err != nil {
		return nil, err
	}

	for _, oi := range ooii {
		orderItems = append(orderItems,
			&pb.OrderItemByProductName{
				Name:     oi.ProductName,
				Quantity: oi.Quantity,
			})
	}

	return orderItems, nil
}
