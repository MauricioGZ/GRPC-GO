package main

import (
	"log"
	"net"

	pb "github.com/MauricioGZ/GRPC-GO/internal/gen"
	"github.com/MauricioGZ/GRPC-GO/internal/server/db"
	"github.com/MauricioGZ/GRPC-GO/internal/server/repository"
	client_service "github.com/MauricioGZ/GRPC-GO/internal/server/service/client"
	restaurant_service "github.com/MauricioGZ/GRPC-GO/internal/server/service/restaurant"

	"github.com/MauricioGZ/GRPC-GO/settings"
	"google.golang.org/grpc"
)

const (
	port string = ":8000"
)

func main() {
	s, err := settings.New()
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.New(*s)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := repository.New(db)

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to start the server %s", err.Error())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterClientServiceServer(grpcServer, client_service.New(r))
	pb.RegisterRestaurantServiceServer(grpcServer, restaurant_service.New(r))

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %s", err.Error())
	}
}
