package service

import (
	"context"

	pb "github.com/MauricioGZ/GRPC-GO/internal/gen"
)

func (s *server) GetMenu(req *pb.GetMenuRequest, stream pb.OrdersService_GetMenuServer) error {

	products, err := s.repo.GetAllProducts(context.Background())
	if err != nil {
		return err
	}
	for _, p := range products {
		res := &pb.GetMenuResponse{
			Product: &pb.Product{
				ProductID: p.ID,
				Name:      p.Name,
				Price:     p.Price,
			},
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
	return nil
}
