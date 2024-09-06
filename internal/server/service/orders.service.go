package service

import (
	"context"
	"math/rand"
	"time"

	pb "github.com/MauricioGZ/GRPC-GO/internal/gen"
)

type server struct {
	pb.UnimplementedOrdersServiceServer
}

// faking a db
var orders []*pb.Order
var products []*pb.Product = []*pb.Product{
	{
		ProductID:   1,
		Price:       9.99,
		Description: "Hamburguesa con papas",
	},
	{
		ProductID:   2,
		Price:       2.99,
		Description: "Malteada",
	},
}

func New() *server {
	return &server{}
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
	for _, p := range products {
		res := &pb.GetMenuResponse{
			Product: p,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
