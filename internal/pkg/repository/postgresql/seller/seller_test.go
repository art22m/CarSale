package seller

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgconn"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSellersRepo_Read(t *testing.T) {
	t.Parallel()

	const (
		query = "SELECT id, name, phone_number, created_at, updated_at FROM seller"
	)

	var ctx = context.Background()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		// Arrange
		s := setUp(t)
		defer s.tearDown()

		s.mockDb.EXPECT().Select(gomock.Any(), gomock.Any(), query).Return(nil)

		// Act
		_, err := s.repo.Read(ctx)

		// Assert
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		t.Parallel()

		// Arrange
		s := setUp(t)
		defer s.tearDown()

		s.mockDb.EXPECT().Select(gomock.Any(), gomock.Any(), query).Return(assert.AnError)

		// Act
		_, err := s.repo.Read(ctx)

		// Assert
		assert.Error(t, err)
	})
}

func TestSellersRepo_Delete(t *testing.T) {
	t.Parallel()

	const (
		query       = "DELETE FROM seller WHERE id = $1"
		id    int64 = 0
	)

	var ctx = context.Background()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		t.Run("deleted", func(t *testing.T) {
			t.Parallel()

			// Arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockDb.EXPECT().Exec(gomock.Any(), query, gomock.Any()).Return(pgconn.CommandTag{'1'}, nil)

			// Act
			ok, err := s.repo.Delete(ctx, id)

			// Assert
			assert.NoError(t, err)
			assert.True(t, ok)
		})

		t.Run("not found", func(t *testing.T) {
			t.Parallel()

			// Arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockDb.EXPECT().Exec(gomock.Any(), query, gomock.Any()).Return(pgconn.CommandTag{}, nil)

			// Act
			ok, err := s.repo.Delete(ctx, id)

			// Assert
			assert.NoError(t, err)
			assert.False(t, ok)
		})

	})

	t.Run("fail", func(t *testing.T) {
		t.Parallel()

		t.Run("internal error", func(t *testing.T) {
			t.Parallel()

			// Arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockDb.EXPECT().Exec(gomock.Any(), query, gomock.Any()).Return(nil, assert.AnError)

			// Act
			_, err := s.repo.Delete(ctx, id)

			// Assert
			assert.Error(t, err)
		})

	})

}

func TestSellersRepo_GetById(t *testing.T) {
	t.Parallel()

	const (
		query       = "SELECT id, name, phone_number, created_at, updated_at FROM seller WHERE id=$1"
		id    int64 = 0
	)

	var ctx = context.Background()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		// Arrange
		s := setUp(t)
		defer s.tearDown()

		s.mockDb.EXPECT().Get(gomock.Any(), gomock.Any(), query, gomock.Any()).Return(nil)

		// Act
		seller, err := s.repo.GetById(ctx, id)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, int64(0), seller.ID)
	})

	t.Run("fail", func(t *testing.T) {
		t.Parallel()

		t.Run("seller not found", func(t *testing.T) {
			t.Parallel()

			// Arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockDb.EXPECT().Get(gomock.Any(), gomock.Any(), query, gomock.Any()).Return(sql.ErrNoRows)

			// Act
			seller, err := s.repo.GetById(ctx, id)

			// Assert
			assert.Error(t, err)
			assert.Nil(t, seller)
		})

		t.Run("internal error", func(t *testing.T) {
			t.Parallel()

			// Arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockDb.EXPECT().Get(gomock.Any(), gomock.Any(), query, gomock.Any()).Return(errors.New("some internal error"))

			// Act
			seller, err := s.repo.GetById(ctx, id)

			// Assert
			assert.Error(t, err)
			assert.NotNil(t, seller)
		})

	})

}
