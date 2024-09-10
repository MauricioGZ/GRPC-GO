package client_service

import (
	"context"
	"time"

	pb "github.com/MauricioGZ/GRPC-GO/internal/gen"
)

func (s *service) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	//check if the passed product's ids do exist
	var totalPrice float32 = 0.0
	for _, order := range req.GetOrderItems() {
		price, err := s.repo.GetProductPriceByID(
			ctx,
			order.GetProductID(),
		)

		if price == nil {
			if err != nil {
				return nil, err
			}
			//TODO: implement an error message if the product does not exist
		}
		totalPrice += *price
	}

	//register the order
	orderID, err := s.repo.InsertOrder(
		ctx,
		req.GetCustomerID(),
		time.Now(),
		"pending",
		totalPrice,
	)

	if err != nil {
		return nil, err
	}

	//register each order item
	for _, order := range req.GetOrderItems() {
		err = s.repo.InsertOrderItem(
			ctx,
			*orderID,
			order.GetProductID(),
			order.GetQuantity(),
		)
		if err != nil {
			return nil, err
		}
	}

	return &pb.CreateOrderResponse{
		OrderID: *orderID,
	}, nil
}
