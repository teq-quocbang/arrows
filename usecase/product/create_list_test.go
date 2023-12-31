package product

import (
	"context"
	"fmt"

	"bou.ke/monkey"
	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"github.com/stretchr/testify/mock"
	"github.com/teq-quocbang/arrows/model"
	"github.com/teq-quocbang/arrows/payload"
	"github.com/teq-quocbang/arrows/repository/product"
	"github.com/teq-quocbang/arrows/util/contexts"
	"github.com/teq-quocbang/arrows/util/token"
)

func (s *TestSuite) TestCreateList() {
	assertion := s.Assertions
	testName := fake.Name()
	testProductType := fake.Car().Type
	testProducerID := uuid.New().String()

	userPrinciple := &token.JWTClaimCustom{
		SessionID: uuid.New(),
		User: token.UserInfo{
			ID:       uuid.New(),
			Email:    "test@teqnological.asia",
			Username: "test_username",
		},
	}
	user := monkey.Patch(contexts.GetUserPrincipleByContext, func(context.Context) *token.JWTClaimCustom {
		return userPrinciple
	})
	defer monkey.Unpatch(user)

	// good case
	{
		// Arrange
		mockProduct := product.NewMockRepository(s.T())
		realType := fake.Car().Type
		realName := fake.Name()
		mockProduct.EXPECT().GetList(s.ctx).ReturnArguments = mock.Arguments{
			[]model.Product{
				{
					ID:          uuid.New(),
					Name:        fmt.Sprintf("%s-1", testName), // same name
					ProductType: testProductType,
					ProducerID:  uuid.MustParse(testProducerID),
				},
				{
					ID:          uuid.New(),
					Name:        fmt.Sprintf("%s-2", testName), // same name
					ProductType: testProductType,
					ProducerID:  uuid.MustParse(testProducerID),
				},
			}, nil}
		mockProduct.EXPECT().CreateList(s.ctx, []model.Product{{
			Name:        realName,
			ProductType: realType,
			ProducerID:  uuid.MustParse(testProducerID),
			Price:       decimal.NewFromInt(15000000),
			CreatedBy:   userPrinciple.User.ID,
			UpdatedBy:   userPrinciple.User.ID,
		}}).ReturnArguments = mock.Arguments{nil}
		u := s.useCase(mockProduct)

		// Act
		reply, err := u.CreateList(s.ctx, &payload.CreateListProductRequest{
			Products: []payload.Product{
				{
					Name:        realName,
					ProductType: realType,
					ProducerID:  testProducerID,
					Price:       "15000000",
				},
				{
					Name:        fmt.Sprintf("%s-1", testName), // same name
					ProductType: testProductType,
					ProducerID:  testProducerID,
					Price:       "15000000",
				},
				{
					Name:        fmt.Sprintf("%s-2", testName), // same name
					ProductType: testProductType,
					ProducerID:  testProducerID,
					Price:       "15000000",
				},
			},
		})

		// Assert
		assertion.NoError(err)
		assertion.Equal(1, len(reply.Product))
	}
}
