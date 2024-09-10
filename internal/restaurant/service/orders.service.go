package service

import (
	"context"
	"io"
	"log"

	pb "github.com/MauricioGZ/GRPC-GO/internal/gen"
)

func (s *serv) GetPendingOrders(ctx context.Context) ([]*pb.GetPendingOrdersResponse, error) {
	var pendingOrders []*pb.GetPendingOrdersResponse
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
			pendingOrders = append(pendingOrders, res)
		}
	}()
	<-done
	return pendingOrders, nil
}
