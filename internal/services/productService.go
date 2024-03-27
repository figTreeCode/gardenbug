package services

import (
	"gardenbug/internal/models"

	"github.com/jackc/pgx/v5"
)

func GetActiveProducts(conn *pgx.Conn) ([]models.Product, error) {
	products, err := models.GetActiveProducts(conn)
	if err != nil {
		return nil, err
	}

	return products, nil
}
