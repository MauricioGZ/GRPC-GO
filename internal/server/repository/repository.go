package repository

import (
	"context"
	"database/sql"

	"github.com/MauricioGZ/GRPC-GO/internal/server/entity"
)

type Repository interface {
	GetAllProducts(ctx context.Context) ([]entity.Product, error)
}

type repo struct {
	db *sql.DB
}

func New(_db *sql.DB) Repository {
	return &repo{
		db: _db,
	}
}
