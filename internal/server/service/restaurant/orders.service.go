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
		res := &pb.GetPendingOrdersResponse{
			OrderID:    order.ID,
			CustomerID: order.CustomerID,
			OrderDate:  timestamppb.New(order.OrderDate),
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
	return nil
}
