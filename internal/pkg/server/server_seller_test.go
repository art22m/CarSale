package server

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"hw5/internal/pkg/repository"
)

func Test_getSeller(t *testing.T) {
	t.Parallel()

	var (
		ctx         = context.Background()
		id    int64 = 1
		name        = "Artem"
		phone       = "111-222"
	)

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		// Arrange
		s := setUp(t)
		defer s.tearDown()

		req, err := http.NewRequest(http.MethodGet, "seller?id=1", bytes.NewReader([]byte{}))
		require.NoError(t, err)

		s.mockSellersRepo.EXPECT().GetById(gomock.Any(), id).Return(&repository.Seller{ID: 1, Name: name, PhoneNumber: phone}, nil)

		// Act
		data, status := s.server.getSeller(ctx, req)

		type sellerResponse struct {
			ID          int64  `json:"id"`
			Name        string `json:"name"`
			PhoneNumber string `json:"phone_number"`
		}

		var unmarshalled sellerResponse
		err = json.Unmarshal(data, &unmarshalled)
		assert.NoError(t, err, "cannot unmarshall response data")

		// Assert
		assert.Equal(t, http.StatusOK, status)
		assert.Equal(t, unmarshalled, sellerResponse{id, name, phone})
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		// Arrange
		tt := []struct {
			name    string
			request *http.Request
			status  int
		}{
			{
				"without_id",
				&http.Request{Method: http.MethodGet, URL: &url.URL{RawQuery: "guid=1"}},
				http.StatusBadRequest,
			},
			{
				"wrong_id",
				&http.Request{Method: http.MethodGet, URL: &url.URL{RawQuery: "id=one"}},
				http.StatusBadRequest,
			},
			{
				"empty",
				&http.Request{Method: http.MethodGet, URL: &url.URL{RawQuery: ""}},
				http.StatusBadRequest,
			},
		}

		for _, tc := range tt {
			tc := tc

			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				s := setUp(t)
				defer s.tearDown()

				// Act
				_, status := s.server.getSeller(ctx, tc.request)

				// Assert
				assert.Equal(t, status, tc.status)
			})
		}
	})
}

func Test_createSeller(t *testing.T) {
	t.Parallel()

	// Arrange
	var (
		ctx         = context.Background()
		id    int64 = 1
		name        = "Artem"
		phone       = "111-222"
	)

	s := setUp(t)
	defer s.tearDown()

	ss := &serverSeller{
		Name:        name,
		PhoneNumber: phone,
	}

	data, err := json.Marshal(ss)
	require.NoError(t, err, "cannot marshall response data")

	req, err := http.NewRequest(http.MethodPost, "seller", bytes.NewReader(data))
	require.NoError(t, err, "cannot create request")

	s.mockSellersRepo.EXPECT().Create(gomock.Any(), &repository.Seller{
		Name:        name,
		PhoneNumber: phone,
	}).Return(id, nil)

	// Act
	newId, status := s.server.createSeller(ctx, req)

	// Assert
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, id, newId)
}

func Test_updateSeller(t *testing.T) {
	t.Parallel()

	// Arrange
	var (
		ctx         = context.Background()
		id    int64 = 1
		name        = "Artem"
		phone       = "111-222"
	)

	s := setUp(t)
	defer s.tearDown()

	ss := &serverSeller{
		ID:          id,
		Name:        name,
		PhoneNumber: phone,
	}

	data, err := json.Marshal(ss)
	require.NoError(t, err, "cannot marshall response data")

	req, err := http.NewRequest(http.MethodPut, "seller", bytes.NewReader(data))
	require.NoError(t, err, "cannot create request")

	s.mockSellersRepo.EXPECT().Update(gomock.Any(), &repository.Seller{
		ID:          id,
		Name:        name,
		PhoneNumber: phone,
	}).Return(true, nil)

	// Act
	ok, status := s.server.updateSeller(ctx, req)

	// Assert
	assert.Equal(t, http.StatusOK, status)
	assert.True(t, ok)
}

func Test_deleteSeller(t *testing.T) {
	t.Parallel()

	// Arrange
	var (
		ctx       = context.Background()
		id  int64 = 1
	)

	s := setUp(t)
	defer s.tearDown()

	req, err := http.NewRequest(http.MethodDelete, "seller?id=1", bytes.NewReader([]byte{}))
	require.NoError(t, err)

	s.mockSellersRepo.EXPECT().Delete(gomock.Any(), id).Return(true, nil)

	// Act
	ok, status := s.server.deleteSeller(ctx, req)

	// Assert
	assert.Equal(t, http.StatusOK, status)
	assert.True(t, ok)
}

func Test_getSellerData(t *testing.T) {
	t.Parallel()

	var (
		id    int64 = 1
		name        = "Artem"
		phone       = "111-222"
	)

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		// Arrange
		reqSeller := &serverSeller{
			ID:          id,
			Name:        name,
			PhoneNumber: phone,
		}

		body, err := json.Marshal(reqSeller)
		require.NoError(t, err, "cannot marshall response data")

		req, err := http.NewRequest(http.MethodGet, "", bytes.NewReader(body))
		require.NoError(t, err, "cannot create request")

		// Act
		res, err := getSellerData(req.Body)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		// Arrange

		body, err := json.Marshal([]byte(`1112`))
		require.NoError(t, err, "cannot marshall response data")

		req, err := http.NewRequest(http.MethodGet, "", bytes.NewReader(body))
		require.NoError(t, err, "cannot create request")

		// Act
		_, err = getSellerData(req.Body)

		// Assert
		assert.Error(t, err)
	})
}
