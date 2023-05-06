package seller

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_database "hw5/internal/pkg/db/mocks"
	"hw5/internal/pkg/repository"
)

type sellerRepoFixture struct {
	ctrl   *gomock.Controller
	repo   repository.SellersRepo
	mockDb *mock_database.MockDatabaseOperations
}

func setUp(t *testing.T) sellerRepoFixture {
	ctrl := gomock.NewController(t)
	mockDb := mock_database.NewMockDatabaseOperations(ctrl)
	repo := NewSellers(mockDb)

	return sellerRepoFixture{repo: repo, ctrl: ctrl, mockDb: mockDb}
}

func (u *sellerRepoFixture) tearDown() {
	u.ctrl.Finish()
}
