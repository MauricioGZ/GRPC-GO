package repository

import (
	"context"
)

const (
	qryInsertOrderItem = `insert into ORDER_ITEMS(
													orderId,
													productId,
													quantity
												)
												values (?,?,?);`
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
