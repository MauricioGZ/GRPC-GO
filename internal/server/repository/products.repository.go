package repository

import (
	"context"
	"database/sql"

	"github.com/MauricioGZ/GRPC-GO/internal/server/entity"
)

const (
	qryGetAllProducts = `	select
													id,
													name,
													price
												from PRODUCTS;`
)

func (r *repo) GetAllProducts(ctx context.Context) ([]entity.Product, error) {
	var product entity.Product
	var products []entity.Product

	rows, err := r.db.QueryContext(
		ctx,
		qryGetAllProducts,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}
