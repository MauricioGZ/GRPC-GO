package client_service

import (
	pb "github.com/MauricioGZ/GRPC-GO/internal/gen"
	"github.com/MauricioGZ/GRPC-GO/internal/server/repository"
)

type service struct {
	repo repository.Repository
	pb.UnimplementedOrdersServiceServer
}

func New(_repo repository.Repository) *service {
	return &service{
		repo: _repo,
	}
}
