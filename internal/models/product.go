package models

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Product struct {
	ProductId            int
	ProductName          string
	ProductDesc          *string
	ProductPrice         *float32
	IsPurchasedInd       bool
	BatchId              int
	Quantity             int
	ProductTypeId        int
	SizeId               *int
	ProductSpecficTypeId int
	DeletedInd           bool
}

func GetActiveProducts(conn *pgx.Conn) ([]Product, error) {
	rows, err := conn.Query(context.Background(), "SELECT * FROM gardenbug.product WHERE deleted_ind = false")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		if err := rows.Scan(
			&product.ProductId,
			&product.ProductName,
			&product.ProductDesc,
			&product.ProductPrice,
			&product.IsPurchasedInd,
			&product.BatchId,
			&product.Quantity,
			&product.ProductTypeId,
			&product.SizeId,
			&product.ProductSpecficTypeId,
			&product.DeletedInd,
		); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
