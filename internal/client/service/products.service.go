package service

import (
	"context"
	"io"
	"log"

	"github.com/MauricioGZ/GRPC-GO/internal/client/model"
	pb "github.com/MauricioGZ/GRPC-GO/internal/gen"
)

func (s *serv) GetAllProducts(ctx context.Context) ([]model.Product, error) {
	var product model.Product
	var products []model.Product
	done := make(chan bool)
	menu, err := s.client.GetMenu(ctx, &pb.GetMenuRequest{})
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			res, err := menu.Recv()
			if err != nil {
				if err == io.EOF {
					done <- true
					return
				}
				log.Fatal(err)
			}
			product = model.Product{
				ProductID: res.GetProduct().GetProductID(),
				Name:      res.GetProduct().GetName(),
				Price:     res.GetProduct().GetPrice(),
			}
			products = append(products, product)
		}
	}()
	<-done
	return products, nil
}
