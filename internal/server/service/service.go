package service

import (
	pb "github.com/MauricioGZ/GRPC-GO/internal/gen"
	"github.com/MauricioGZ/GRPC-GO/internal/server/repository"
)

type server struct {
	repo repository.Repository
	pb.UnimplementedOrdersServiceServer
}

func New(_repo repository.Repository) *server {
	return &server{
		repo: _repo,
	}
}
