//go:build integration

package tests

import (
	"context"
	"github.com/stretchr/testify/assert"
	"hw5/internal/pkg/repository"
	"hw5/internal/pkg/repository/postgresql/seller"
	"testing"
)

func TestSellersRepo_Delete(t *testing.T) {
	TestDatabase.SetUp(t)
	defer TestDatabase.TearDown(t)

	// Arrange
	sellersRepo := seller.NewSellers(TestDatabase.DB)

	const (
		name  = "ivan"
		phone = "111-222"
	)

	t.Run("success", func(t *testing.T) {
		// Act
		id, err := sellersRepo.Create(context.Background(), &repository.Seller{
			Name:        name,
			PhoneNumber: phone,
		})

		ok, err := sellersRepo.Delete(context.Background(), id)

		// Assert
		assert.NoError(t, err)
		assert.True(t, ok)
	})

	t.Run("failure", func(t *testing.T) {
		// Act

		// Нет такого продавца
		ok, err := sellersRepo.Delete(context.Background(), -1)

		// Assert
		assert.NoError(t, err)
		assert.False(t, ok)
	})
}

func TestSellersRepo_Update(t *testing.T) {
	TestDatabase.SetUp(t)
	defer TestDatabase.TearDown(t)

	// Arrange
	sellersRepo := seller.NewSellers(TestDatabase.DB)

	const (
		name     = "ivan"
		newPhone = "111-222"
		oldPhone = "222-333"
	)

	t.Run("success", func(t *testing.T) {
		// Act
		id, err := sellersRepo.Create(context.Background(), &repository.Seller{
			Name:        name,
			PhoneNumber: oldPhone,
		})

		ok, err := sellersRepo.Update(context.Background(), &repository.Seller{
			ID:          id,
			Name:        name,
			PhoneNumber: newPhone,
		})

		// Assert
		assert.NoError(t, err)
		assert.True(t, ok)
	})

	t.Run("failure", func(t *testing.T) {
		// Act

		// Нет такого продавца
		ok, err := sellersRepo.Update(context.Background(), &repository.Seller{
			ID:          -1,
			Name:        "ivan",
			PhoneNumber: "111-222",
		})

		// Assert
		assert.NoError(t, err)
		assert.False(t, ok)
	})
}

func TestSellersRepo_Create(t *testing.T) {
	TestDatabase.SetUp(t)
	defer TestDatabase.TearDown(t)

	// Arrange
	sellersRepo := seller.NewSellers(TestDatabase.DB)

	t.Run("success", func(t *testing.T) {
		// Act
		_, err := sellersRepo.Create(context.Background(), &repository.Seller{
			Name:        "ivan",
			PhoneNumber: "111-222",
		})

		// Assert
		assert.NoError(t, err)
	})
}

func TestSellersRepo_Read(t *testing.T) {
	TestDatabase.SetUp(t)
	defer TestDatabase.TearDown(t)

	// Arrange
	sellersRepo := seller.NewSellers(TestDatabase.DB)

	const (
		name  = "ivan"
		phone = "111-222"
	)

	t.Run("success", func(t *testing.T) {
		// Act
		id, err := sellersRepo.Create(context.Background(), &repository.Seller{
			Name:        name,
			PhoneNumber: phone,
		})

		sellers, err := sellersRepo.Read(context.Background())

		// Assert
		contains := false
		for _, sl := range sellers {
			if sl.ID == id {
				contains = true
				break
			}
		}

		assert.NoError(t, err)
		assert.True(t, contains)
	})

	t.Run("failure", func(t *testing.T) {
		// Act
		sellers, err := sellersRepo.Read(context.Background())

		// Assert
		contains := false
		for _, sl := range sellers {
			if sl.ID == -1 {
				contains = true
				break
			}
		}

		assert.NoError(t, err)
		assert.False(t, contains)
	})
}

func TestSellersRepo_GetById(t *testing.T) {
	TestDatabase.SetUp(t)
	defer TestDatabase.TearDown(t)

	// Arrange
	sellersRepo := seller.NewSellers(TestDatabase.DB)

	const (
		name  = "ivan"
		phone = "111-222"
	)

	t.Run("success", func(t *testing.T) {
		// Act
		id, err := sellersRepo.Create(context.Background(), &repository.Seller{
			Name:        name,
			PhoneNumber: phone,
		})

		s, err := sellersRepo.GetById(context.Background(), id)

		// Assert

		assert.NoError(t, err)
		assert.Equal(t, s.Name, name)
		assert.Equal(t, s.PhoneNumber, phone)
	})

	t.Run("failure", func(t *testing.T) {
		// Act
		_, err := sellersRepo.GetById(context.Background(), -1)

		// Assert
		assert.Error(t, err)
	})
}
