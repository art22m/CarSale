package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"hw5/internal/pkg/db"
	"strings"
	"sync"
	"testing"
)

type TestDatabase struct {
	sync.Mutex
	DB *db.Database
}

const (
	host     = "localhost"
	port     = 5432
	user     = "test"
	password = "test"
	dbname   = "gojhw5"
)

func NewTestDatabase() (*TestDatabase, error) {
	pool, err := pgxpool.Connect(context.Background(), generateDsn())
	if err != nil {
		return nil, err
	}

	return &TestDatabase{DB: db.NewDatabase(pool)}, nil
}

func generateDsn() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}

func (d *TestDatabase) SetUp(t *testing.T) {
	t.Helper()
	ctx := context.Background()
	d.Lock()
	d.Truncate(ctx)
}

func (d *TestDatabase) TearDown(t *testing.T) {
	defer d.Unlock()
	d.Truncate(context.Background())
}

func (d *TestDatabase) Truncate(ctx context.Context) {
	var tables []string
	err := d.DB.Select(ctx, &tables, "SELECT table_name FROM information_schema.tables WHERE table_schema='public' AND table_type='BASE TABLE' AND table_name != 'goose_db_version'")
	if err != nil {
		panic(err)
	}

	if len(tables) == 0 {
		panic("run migrations -- no tables")
	}

	q := fmt.Sprintf("Truncate table %s", strings.Join(tables, ","))
	if _, err = d.DB.Exec(ctx, q); err != nil {
		panic(err)
	}
}
