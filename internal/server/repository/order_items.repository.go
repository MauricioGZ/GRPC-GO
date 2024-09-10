package repository

import (
	"context"
	"database/sql"

	"github.com/MauricioGZ/GRPC-GO/internal/server/entity"
)

const (
	qryInsertOrderItem = `insert into ORDER_ITEMS(
													orderId,
													productId,
													quantity
												)
												values (?,?,?);`
	qryGetOrderItemsByOrderID = `	select
																	PRODUCTS.name,
																	ORDER_ITEMS.quantity
																from ORDER_ITEMS
																join PRODUCTS on ORDER_ITEMS.productId = PRODUCTS.id
																where orderId = ?;`
)

func (r *repo) InsertOrderItem(ctx context.Context, orderID, productID, quantity uint32) error {
	_, err := r.db.ExecContext(
		ctx,
		qryInsertOrderItem,
		orderID,
		productID,
		quantity,
	)

	return err
}

func (r *repo) GetOrderItemsByOrderID(ctx context.Context, orderID uint32) ([]entity.OrderItem, error) {
	var orderItem entity.OrderItem
	var orderItems []entity.OrderItem

	rows, err := r.db.QueryContext(
		ctx,
		qryGetOrderItemsByOrderID,
		orderID,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&orderItem.ProductName,
			&orderItem.Quantity,
		)

		if err != nil {
			return nil, err
		}

		orderItems = append(orderItems, orderItem)
	}
	return orderItems, nil
}
