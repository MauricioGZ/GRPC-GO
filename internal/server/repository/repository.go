package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/MauricioGZ/GRPC-GO/internal/server/entity"
)

type Repository interface {
	//products respository
	GetProductPriceByID(ctx context.Context, producID uint32) (*float32, error)
	GetAllProducts(ctx context.Context) ([]entity.Product, error)
	//orders repository
	GetOrderByID(ctx context.Context, orderID uint32) (*entity.Order, error)
	InsertOrder(ctx context.Context, customerID uint32, orderDate time.Time, status string, totalPrice float32) (*uint32, error)
	UpdateOrderStatus(ctx context.Context, status string, orderID uint32) error
	GetPendingOrders(ctx context.Context) ([]entity.Order, error)
	//order items repository
	InsertOrderItem(ctx context.Context, orderID, productID, quantity uint32) error
	GetOrderItemsByOrderID(ctx context.Context, orderID uint32) ([]entity.OrderItem, error)
}

type repo struct {
	db *sql.DB
}

func New(_db *sql.DB) Repository {
	return &repo{
		db: _db,
	}
}
