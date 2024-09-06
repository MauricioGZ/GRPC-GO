package api

import (
	"log"
	"net/http"

	pb "github.com/MauricioGZ/GRPC-GO/internal/gen"
)

type api struct {
	client pb.OrdersServiceClient
}

func New(_client pb.OrdersServiceClient) *api {
	return &api{
		client: _client,
	}
}

func (a *api) Run() error {
	router := http.NewServeMux()
	a.registerRoutes(router)

	log.Println("Starting server on ", "9000")
	return http.ListenAndServe(":9000", router)
}
