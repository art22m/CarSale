package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "test"
	password = "test"
	dbname   = "gojhw5"
)

func CreateDatabase(ctx context.Context) (*Database, error) {
	pool, err := pgxpool.Connect(ctx, generateDsn())
	if err != nil {
		return nil, err
	}

	return NewDatabase(pool), nil
}

func generateDsn() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}
