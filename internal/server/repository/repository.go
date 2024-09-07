package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/MauricioGZ/GRPC-GO/internal/server/entity"
)

type Repository interface {
	GetProductPriceByID(ctx context.Context, producID uint32) (*float32, error)
	GetAllProducts(ctx context.Context) ([]entity.Product, error)
	InsertOrder(ctx context.Context, customerID uint32, orderDate time.Time, status string, totalPrice float32) (*uint32, error)
	InsertOrderItem(ctx context.Context, orderID, productID, quantity uint32) error
}

type repo struct {
	db *sql.DB
}

func New(_db *sql.DB) Repository {
	return &repo{
		db: _db,
	}
}
