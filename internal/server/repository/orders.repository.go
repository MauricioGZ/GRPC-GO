package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/MauricioGZ/GRPC-GO/internal/server/entity"
)

const (
	qryGetOrderByID = ` select
												id,
												customerid,
												orderDate,
												status,
												totalPrice
											from ORDERS
											where id = ?;`
	qryInsertOrder = `insert into ORDERS(
											customerId,
											orderDate,
											status,
											totalPrice
											)
										values (?,?,?,?)`
	qryLastInsertID     = `	select LAST_INSERT_ID();`
	qryGetPendingOrders = `	select
														id,
														customerId,
														orderDate
													from ORDERS
													where status = "pending";`
	qryUpdateOrderStatus = `	update ORDERS
														set
															status = ?
														where id = ?;`
)

func (r *repo) GetOrderByID(ctx context.Context, orderID uint32) (*entity.Order, error) {
	var order entity.Order
	row := r.db.QueryRowContext(
		ctx,
		qryGetOrderByID,
		orderID,
	)

	err := row.Scan(
		&order.ID,
		&order.CustomerID,
		&order.OrderDate,
		&order.Status,
		&order.TotalPrice,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &order, nil
}

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

func (r *repo) GetPendingOrders(ctx context.Context) ([]entity.Order, error) {
	var order entity.Order
	var orders []entity.Order
	rows, err := r.db.QueryContext(
		ctx,
		qryGetPendingOrders,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&order.ID,
			&order.CustomerID,
			&order.OrderDate,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (r *repo) UpdateOrderStatus(ctx context.Context, status string, orderID uint32) error {
	_, err := r.db.ExecContext(
		ctx,
		qryUpdateOrderStatus,
		status,
		orderID,
	)
	return err
}
