package service

import (
	"context"
	"math/rand"
	"time"

	pb "github.com/MauricioGZ/GRPC-GO/internal/gen"
	"github.com/MauricioGZ/GRPC-GO/internal/server/repository"
)

type server struct {
	repo repository.Repository
	pb.UnimplementedOrdersServiceServer
}

var orders []*pb.Order

func New(_repo repository.Repository) *server {
	return &server{
		repo: _repo,
	}
}

func (s *server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	var order pb.Order = pb.Order{
		OrderID:    generateID(),
		CustomerID: req.CustomerID,
		ProductID:  req.ProductID,
		Quantity:   req.Quantity,
	}

	orders = append(orders, &order)

	return &pb.CreateOrderResponse{
		Orders: orders,
	}, nil

}

func generateID() uint32 {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	return r.Uint32()
}

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
