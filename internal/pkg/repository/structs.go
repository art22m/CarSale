package repository

import (
	"database/sql"
	"time"
)

type CarSale struct {
	ID       int64  `db:"id"`
	Brand    string `db:"brand"`
	Model    string `db:"model"`
	SellerID int64  `db:"seller_id"`

	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type Seller struct {
	ID          int64  `db:"id"`
	Name        string `db:"name"`
	PhoneNumber string `db:"phone_number"`

	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
