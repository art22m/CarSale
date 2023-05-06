//go:build integration

package tests

import (
	"context"
	"github.com/stretchr/testify/assert"
	"hw5/internal/pkg/repository"
	carSale "hw5/internal/pkg/repository/postgresql/car_sale"
	"testing"
)

func TestCarSalesRepo_Delete(t *testing.T) {
	TestDatabase.SetUp(t)
	defer TestDatabase.TearDown(t)

	// Arrange
	carSalesRepo := carSale.NewCarSales(TestDatabase.DB)

	const (
		brand          = "Lada"
		model          = "Calina"
		sellerID int64 = 1
	)

	t.Run("success", func(t *testing.T) {
		// Act
		id, err := carSalesRepo.Create(context.Background(), &repository.CarSale{
			Brand:    brand,
			Model:    model,
			SellerID: sellerID,
		})

		ok, err := carSalesRepo.Delete(context.Background(), id)

		// Assert
		assert.NoError(t, err)
		assert.True(t, ok)
	})

	t.Run("failure", func(t *testing.T) {
		// Act

		// Нет такого продавца
		ok, err := carSalesRepo.Delete(context.Background(), -1)

		// Assert
		assert.NoError(t, err)
		assert.False(t, ok)
	})
}

func TestCarSalesRepo_GetById(t *testing.T) {
	TestDatabase.SetUp(t)
	defer TestDatabase.TearDown(t)

	// Arrange
	carSalesRepo := carSale.NewCarSales(TestDatabase.DB)

	const (
		brand          = "Lada"
		model          = "Calina"
		sellerID int64 = 1
	)

	t.Run("success", func(t *testing.T) {
		// Act
		id, err := carSalesRepo.Create(context.Background(), &repository.CarSale{
			Brand:    brand,
			Model:    model,
			SellerID: sellerID,
		})

		cs, err := carSalesRepo.GetById(context.Background(), id)

		// Assert

		assert.NoError(t, err)
		assert.Equal(t, cs.Brand, brand)
		assert.Equal(t, cs.Model, model)
		assert.Equal(t, cs.SellerID, sellerID)
	})

	t.Run("failure", func(t *testing.T) {
		// Act
		_, err := carSalesRepo.GetById(context.Background(), -1)

		// Assert
		assert.Error(t, err)
	})
}
