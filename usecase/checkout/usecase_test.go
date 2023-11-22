package checkout

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/teq-quocbang/arrows/config"
	"github.com/teq-quocbang/arrows/repository/checkout"
	"github.com/teq-quocbang/arrows/repository/product"
	"github.com/teq-quocbang/arrows/repository/storage"
)

type TestSuite struct {
	suite.Suite

	ctx context.Context

	useCase func(*checkout.MockRepository, *storage.MockRepository, *product.MockRepository) UseCase
}

func (suite *TestSuite) SetupTest() {
	suite.ctx = context.Background()

	suite.useCase = NewTestUseCase
}

func TestUseCaseAuth(t *testing.T) {
	t.Parallel()
	suite.Run(t, &TestSuite{})
}

func NewTestUseCase(checkout *checkout.MockRepository, storage *storage.MockRepository, product *product.MockRepository) UseCase {
	return UseCase{
		Checkout: checkout,
		Storage:  storage,
		Product:  product,
		Config:   config.GetConfig(),
	}
}
