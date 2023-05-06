package server

import (
	"github.com/golang/mock/gomock"
	mock_repository "hw5/internal/pkg/repository/mocks"
	"testing"
)

type serverFixture struct {
	ctrl             *gomock.Controller
	mockSellersRepo  *mock_repository.MockSellersRepo
	mockCarSalesRepo *mock_repository.MockCarSalesRepo
	server           server
}

func setUp(t *testing.T) serverFixture {
	ctrl := gomock.NewController(t)
	mockSellersRepo := mock_repository.NewMockSellersRepo(ctrl)
	mockCarSalesRepo := mock_repository.NewMockCarSalesRepo(ctrl)

	server := server{
		sellersRepo: mockSellersRepo,
		carSaleRepo: mockCarSalesRepo,
	}

	return serverFixture{
		ctrl:             ctrl,
		mockSellersRepo:  mockSellersRepo,
		mockCarSalesRepo: mockCarSalesRepo,
		server:           server,
	}
}

func (sf *serverFixture) tearDown() {
	sf.ctrl.Finish()
}
