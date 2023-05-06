//go:build integration

package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"hw5/internal/pkg/repository"
	"hw5/internal/pkg/repository/postgresql/seller"
	"io"
	"net/http"
	"testing"
)

func Test_createSeller(t *testing.T) {
	// Arrange
	var (
		name  = "Artem"
		phone = "111-222"
	)

	ss := struct {
		ID          int64  `json:"id"`
		Name        string `json:"name"`
		PhoneNumber string `json:"phone_number"`
	}{
		Name:        name,
		PhoneNumber: phone,
	}

	data, err := json.Marshal(ss)
	require.NoError(t, err, "cannot marshall response data")

	req, err := http.NewRequest(http.MethodPost, TestURL+"/seller", bytes.NewReader(data))
	require.NoError(t, err, "cannot create request")

	// Act
	resp, err := http.DefaultClient.Do(req)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func Test_getSeller(t *testing.T) {
	TestDatabase.SetUp(t)
	defer TestDatabase.TearDown(t)

	// Arrange
	sellersRepo := seller.NewSellers(TestDatabase.DB)

	const (
		name  = "ivan"
		phone = "111-222"
	)

	// Act
	id, err := sellersRepo.Create(context.Background(), &repository.Seller{
		Name:        name,
		PhoneNumber: phone,
	})
	require.NoError(t, err)

	url := fmt.Sprintf("%s/seller?id=%v", TestURL, id)
	req, err := http.NewRequest(http.MethodGet, url, bytes.NewReader([]byte{}))
	require.NoError(t, err, "cannot create request")

	resp, err := http.DefaultClient.Do(req)

	// Assert
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	data := struct {
		ID          int64  `json:"id"`
		Name        string `json:"name"`
		PhoneNumber string `json:"phone_number"`
	}{}
	err = json.Unmarshal(body, &data)

	assert.NoError(t, err)
	assert.Equal(t, data.ID, id)
	assert.Equal(t, data.Name, name)
	assert.Equal(t, data.PhoneNumber, phone)
}
