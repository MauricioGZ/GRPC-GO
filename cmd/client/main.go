package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/MauricioGZ/GRPC-GO/internal/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port string = ":8000"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Did not connect: %s", err.Error())
	}

	defer conn.Close()

	client := pb.NewOrdersServiceClient(conn)
	menu, err := client.GetMenu(context.Background(), &pb.GetMenuRequest{})

	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)

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
			fmt.Println(res.Product)
		}
	}()
	<-done
}
