package main

import (
	"log"

	"github.com/MauricioGZ/GRPC-GO/internal/client/api"
	"github.com/MauricioGZ/GRPC-GO/internal/client/service"
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

	s := service.New(client)
	a := api.New(s)

	err = a.Run()
	if err != nil {
		log.Fatal(err)
	}
}
