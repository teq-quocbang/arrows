package checkout

import (
	"context"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	"github.com/teq-quocbang/store/delivery/http/auth"
	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/repository/checkout"
	"github.com/teq-quocbang/store/repository/product"
	"github.com/teq-quocbang/store/repository/storage"
	"github.com/teq-quocbang/store/util/token"
)

func (s *TestSuite) TestAddToCart() {
	assertion := s.Assertions
	testProductID := uuid.New()
	testQty := 1

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
		mockStorage := storage.NewMockRepository(s.T())
		mockStorage.EXPECT().GetInventoryQty(ctx, testProductID).ReturnArguments = mock.Arguments{10, nil}
		mockCheckout.EXPECT().UpsertCart(ctx, &model.Cart{
			AccountID: userPrinciple.User.ID,
			ProductID: testProductID,
			Qty:       int64(testQty),
		}).ReturnArguments = mock.Arguments{nil}
		req := &payload.AddToCartRequest{
			ProductID: testProductID.String(),
			Qty:       int64(testQty),
		}
		u := s.useCase(mockCheckout, mockStorage, product.NewMockRepository(s.T()))

		// Act
		reply, err := u.AddToCard(ctx, req)

		// Assert
		assertion.NoError(err)
		assertion.Equal(testProductID, reply.Cart.ProductID)
		assertion.Equal(int64(testQty), reply.Cart.Qty)
	}

	// bad case
	{
		// Arrange
		u := s.useCase(checkout.NewMockRepository(s.T()), storage.NewMockRepository(s.T()), product.NewMockRepository(s.T()))

		// Act
		_, err := u.AddToCard(ctx, &payload.AddToCartRequest{})

		// Assert
		assertion.Error(err)
	}
}
