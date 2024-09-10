package service

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/MauricioGZ/GRPC-GO/internal/gen"
)

func (s *serv) GetPendingOrders(ctx context.Context) error {
	done := make(chan bool)
	orders, err := s.client.GetPendingOrders(ctx, &pb.GetPendingOrdersRequest{})
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			res, err := orders.Recv()
			if err != nil {
				if err == io.EOF {
					done <- true
					return
				}
				log.Fatal(err)
			}
			fmt.Println(res)
		}
	}()
	<-done
	return nil
}
