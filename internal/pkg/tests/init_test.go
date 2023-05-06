//go:build integration

package tests

import (
	"context"
	carSale "hw5/internal/pkg/repository/postgresql/car_sale"
	"hw5/internal/pkg/repository/postgresql/seller"
	"hw5/internal/pkg/server"
	"hw5/internal/pkg/tests/postgresql"
	"log"
	"net/http"
)

var (
	TestServerPort = ":9002"
	TestHost       = "localhost"
	TestURL        = "http://" + TestHost + TestServerPort

	TestDatabase *postgresql.TestDatabase
)

func init() {
	database, err := postgresql.NewTestDatabase()
	if err != nil {
		log.Fatal("error while init test db: ", err)
	}
	TestDatabase = database

	serverMux := server.CreateServer(context.Background(), seller.NewSellers(TestDatabase.DB), carSale.NewCarSales(TestDatabase.DB))
	go func() {
		if err := http.ListenAndServe(TestServerPort, serverMux); err != nil {
			log.Fatal(err)
		}
	}()
}
