//go:generate mockgen -source ./repository.go -destination=./mocks/repository.go -package=mock_repository
package repository

import "context"

type SellersRepo interface {
	Create(ctx context.Context, user *Seller) (int64, error)
	Read(ctx context.Context) ([]*Seller, error)
	Update(ctx context.Context, user *Seller) (bool, error)
	Delete(ctx context.Context, id int64) (bool, error)
	GetById(ctx context.Context, id int64) (*Seller, error)
}

type CarSalesRepo interface {
	Create(ctx context.Context, user *CarSale) (int64, error)
	Read(ctx context.Context) ([]*CarSale, error)
	Update(ctx context.Context, user *CarSale) (bool, error)
	Delete(ctx context.Context, id int64) (bool, error)
	GetById(ctx context.Context, id int64) (*CarSale, error)
}
