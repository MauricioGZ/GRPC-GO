package repository

import (
	"context"
	"time"
)

const (
	qryInsertOrder = `insert into ORDERS(
											customerId,
											orderDate,
											status,
											totalPrice
											)
										values (?,?,?,?)`
	qryLastInsertID = `select LAST_INSERT_ID();`
)

func (r *repo) InsertOrder(ctx context.Context, customerID uint32, orderDate time.Time, status string, totalPrice float32) (*uint32, error) {
	var orderID uint32
	_, err := r.db.ExecContext(
		ctx,
		qryInsertOrder,
		customerID,
		orderDate.UTC(),
		status,
		totalPrice,
	)

	if err != nil {
		return nil, err
	}

	err = r.db.QueryRowContext(
		ctx,
		qryLastInsertID,
	).Scan(
		&orderID,
	)

	if err != nil {
		return nil, err
	}

	return &orderID, err
}
