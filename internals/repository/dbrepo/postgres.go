package dbrepo

import (
	"context"
	"time"

	"github.com/ishanshre/goFiberApiExample/internals/models"
)

func (m *postgresDBRepo) GetAllProducts(limit, offset int) ([]*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `select * from products limit $1 offset $2`
	products := []*models.Product{}
	rows, err := m.DB.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		product := new(models.Product)
		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Body,
			&product.Stock,
			&product.CreatedAt,
			&product.UpdatedAt,
		); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, err
}

func (m *postgresDBRepo) InsertProduct(p *models.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `insert into products (name, body, stock, created_at, updated_at) values ($1,$2,$3,$4,$5)`
	_, err := m.DB.ExecContext(ctx, stmt, p.Name, p.Body, p.Stock, p.CreatedAt, p.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
