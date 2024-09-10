package main

import (
	"log"

	pb "github.com/MauricioGZ/GRPC-GO/internal/gen"
	as "github.com/MauricioGZ/GRPC-GO/internal/restaurant/api"
	sr "github.com/MauricioGZ/GRPC-GO/internal/restaurant/service"
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

	client := pb.NewRestaurantServiceClient(conn)

	s := sr.New(client)
	a := as.New(s)

	err = a.Run()
	if err != nil {
		log.Fatal(err)
	}
}
