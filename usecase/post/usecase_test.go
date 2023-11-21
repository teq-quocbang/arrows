package post

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/teq-quocbang/store/config"
)

type TestSuite struct {
	suite.Suite

	ctx context.Context

	useCase func() UseCase
}

func (suite *TestSuite) SetupTest() {
	suite.ctx = context.Background()

	suite.useCase = NewTestUseCase
}

func TestUseCaseAuth(t *testing.T) {
	t.Parallel()
	suite.Run(t, &TestSuite{})
}

func NewTestUseCase() UseCase {
	return UseCase{
		Config: config.GetConfig(),
	}
}
