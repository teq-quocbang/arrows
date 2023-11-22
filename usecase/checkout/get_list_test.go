package checkout

import (
	"context"

	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/mock"

	"github.com/teq-quocbang/arrows/delivery/http/auth"
	"github.com/teq-quocbang/arrows/model"
	"github.com/teq-quocbang/arrows/repository/checkout"
	"github.com/teq-quocbang/arrows/repository/product"
	"github.com/teq-quocbang/arrows/repository/storage"
	"github.com/teq-quocbang/arrows/util/token"
)

func (s *TestSuite) TestGetListCart() {
	assertion := s.Assertions
	testProductID := uuid.New()

	userPrinciple := &token.JWTClaimCustom{
		SessionID: uuid.New(),
		User: token.UserInfo{
			ID:       uuid.New(),
			Email:    "test@teqnological.asia",
			Username: "test_username",
		},
	}
	ctx := context.WithValue(s.ctx, auth.UserPrincipleKey, userPrinciple)

	// good case
	{
		// Arrange
		mockCheckout := checkout.NewMockRepository(s.T())
		mockProduct := product.NewMockRepository(s.T())

		mockCheckout.EXPECT().GetListCart(ctx, userPrinciple.User.ID).ReturnArguments = mock.Arguments{
			[]model.Cart{
				{
					AccountID: userPrinciple.User.ID,
					ProductID: testProductID,
					Qty:       int64(fake.Uint8()),
				},
			}, nil}
		mockProduct.EXPECT().GetByID(ctx, testProductID).ReturnArguments = mock.Arguments{model.Product{
			ID:          uuid.New(),
			Name:        fake.Name(),
			ProductType: fake.Car().Type,
			ProducerID:  testProductID,
			Price:       decimal.NewFromInt(150000000),
		}, nil}
		u := s.useCase(mockCheckout, storage.NewMockRepository(s.T()), mockProduct)

		// Act
		reply, err := u.GetListCart(ctx)

		// Assert
		assertion.NoError(err)
		assertion.NotNil(reply)
	}
}
