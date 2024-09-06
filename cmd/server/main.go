package main

import (
	"log"
	"net"

	pb "github.com/MauricioGZ/GRPC-GO/internal/gen"
	serverService "github.com/MauricioGZ/GRPC-GO/internal/server/service"
	"google.golang.org/grpc"
)

const (
	port string = ":8000"
)

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to start the server %s", err.Error())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterOrdersServiceServer(grpcServer, serverService.New())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %s", err.Error())
	}
}
