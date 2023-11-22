package checkout

import (
	"context"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/teq-quocbang/arrows/delivery/http/auth"
	"github.com/teq-quocbang/arrows/model"
	"github.com/teq-quocbang/arrows/payload"
	"github.com/teq-quocbang/arrows/repository/checkout"
	"github.com/teq-quocbang/arrows/repository/product"
	"github.com/teq-quocbang/arrows/repository/storage"
	"github.com/teq-quocbang/arrows/util/token"
)

func (s *TestSuite) TestRemoveFromCart() {
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
		mockCheckout.EXPECT().GetCartByConstraint(ctx, userPrinciple.User.ID, testProductID).ReturnArguments = mock.Arguments{
			model.Cart{
				AccountID: userPrinciple.User.ID,
				ProductID: testProductID,
				Qty:       10,
			}, nil}
		mockCheckout.EXPECT().RemoveFromCart(ctx, userPrinciple.User.ID, testProductID, int64(2)).ReturnArguments = mock.Arguments{nil}
		req := &payload.RemoveFormCartRequest{
			ProductID: testProductID.String(),
			Qty:       2,
		}
		u := s.useCase(mockCheckout, storage.NewMockRepository(s.T()), product.NewMockRepository(s.T()))

		// Act
		reply, err := u.RemoveFromCart(ctx, req)

		// Assert
		assertion.NoError(err)
		assertion.Equal(int64(8), reply.Cart.Qty)
	}
}
