package service

import (
	"context"
	"io"
	"log"

	pb "github.com/MauricioGZ/GRPC-GO/internal/gen"
)

func (s *serv) GetAllProducts(ctx context.Context) ([]*pb.Product, error) {
	var listOfProducts pb.ListOfProducts
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
			listOfProducts.Products = append(listOfProducts.GetProducts(), res.GetProduct())
		}
	}()
	<-done
	return listOfProducts.GetProducts(), nil
}
