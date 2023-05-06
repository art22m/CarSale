package car_sale

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"hw5/internal/pkg/db"
	"hw5/internal/pkg/repository"
)

type CarSalesRepo struct {
	db db.DatabaseOperations
}

func NewCarSales(db db.DatabaseOperations) *CarSalesRepo {
	return &CarSalesRepo{db: db}
}

func (r *CarSalesRepo) Create(ctx context.Context, carSale *repository.CarSale) (int64, error) {
	var id int64
	err := r.db.ExecQueryRow(ctx, "INSERT INTO car_sale(brand, model, seller_id) VALUES ($1, $2, $3) RETURNING id", carSale.Brand, carSale.Model, carSale.SellerID).Scan(&id)
	return id, err
}

func (r *CarSalesRepo) Read(ctx context.Context) ([]*repository.CarSale, error) {
	carSales := make([]*repository.CarSale, 0)
	err := r.db.Select(ctx, &carSales, "SELECT id, brand, model, seller_id, created_at, updated_at FROM car_sale")
	return carSales, err
}

func (r *CarSalesRepo) Update(ctx context.Context, user *repository.CarSale) (bool, error) {
	result, err := r.db.Exec(ctx, "UPDATE car_sale SET brand = $1, model = $2, seller_id = $3, updated_at = now() WHERE id = $4", user.Brand, user.Model, user.SellerID, user.ID)
	return result.RowsAffected() > 0, err
}

func (r *CarSalesRepo) Delete(ctx context.Context, id int64) (bool, error) {
	result, err := r.db.Exec(ctx,
		"DELETE FROM car_sale WHERE id = $1", id)
	return result.RowsAffected() > 0, err
}

func (r *CarSalesRepo) GetById(ctx context.Context, id int64) (*repository.CarSale, error) {
	var u repository.CarSale
	err := r.db.Get(ctx, &u, "SELECT id, brand, model, seller_id, created_at, updated_at FROM car_sale WHERE id=$1", id)
	if err == sql.ErrNoRows {
		return nil, errors.New(fmt.Sprintf("car sale with id = %v not found", id))
	}

	return &u, err
}
