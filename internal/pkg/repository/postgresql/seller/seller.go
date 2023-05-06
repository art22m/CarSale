package seller

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"

	"hw5/internal/pkg/db"
	"hw5/internal/pkg/repository"
)

type SellersRepo struct {
	db db.DatabaseOperations
}

func NewSellers(db db.DatabaseOperations) *SellersRepo {
	return &SellersRepo{db: db}
}

func (r *SellersRepo) Create(ctx context.Context, seller *repository.Seller) (int64, error) {
	tr := otel.Tracer("CreateSeller")
	ctx, span := tr.Start(ctx, "repository layer")
	span.SetAttributes(attribute.Key("params").String(fmt.Sprintf("%+v", seller)))
	defer span.End()

	var id int64
	err := r.db.ExecQueryRow(ctx, `INSERT INTO seller(name, phone_number) VALUES ($1, $2) RETURNING id`, seller.Name, seller.PhoneNumber).Scan(&id)
	return id, err
}

func (r *SellersRepo) Read(ctx context.Context) ([]*repository.Seller, error) {
	sellers := make([]*repository.Seller, 0)
	err := r.db.Select(ctx, &sellers, "SELECT id, name, phone_number, created_at, updated_at FROM seller")
	return sellers, err
}

func (r *SellersRepo) Update(ctx context.Context, user *repository.Seller) (bool, error) {
	result, err := r.db.Exec(ctx,
		"UPDATE seller SET name = $1, phone_number = $2, updated_at = now() WHERE id = $3", user.Name, user.PhoneNumber, user.ID)
	return result.RowsAffected() > 0, err
}

func (r *SellersRepo) Delete(ctx context.Context, id int64) (bool, error) {
	result, err := r.db.Exec(ctx,
		"DELETE FROM seller WHERE id = $1", id)
	return result.RowsAffected() > 0, err
}

func (r *SellersRepo) GetById(ctx context.Context, id int64) (*repository.Seller, error) {
	var u repository.Seller
	err := r.db.Get(ctx, &u, "SELECT id, name, phone_number, created_at, updated_at FROM seller WHERE id=$1", id)
	if err == sql.ErrNoRows {
		return nil, errors.New(fmt.Sprintf("seller with id = %v not found", id))
	}

	return &u, err
}
