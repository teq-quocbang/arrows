package post

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/teq-quocbang/arrows/config"
	"github.com/teq-quocbang/arrows/repository/post"
)

type TestSuite struct {
	suite.Suite

	ctx context.Context

	useCase func(*post.MockRepository) UseCase
}

func (suite *TestSuite) SetupTest() {
	suite.ctx = context.Background()

	suite.useCase = NewTestUseCase
}

func TestUseCaseAuth(t *testing.T) {
	t.Parallel()
	suite.Run(t, &TestSuite{})
}

func NewTestUseCase(post *post.MockRepository) UseCase {
	return UseCase{
		Post:   post,
		Config: config.GetConfig(),
	}
}
