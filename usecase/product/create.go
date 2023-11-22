package product

import (
	"context"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"github.com/teq-quocbang/arrows/model"
	"github.com/teq-quocbang/arrows/payload"
	"github.com/teq-quocbang/arrows/presenter"
	"github.com/teq-quocbang/arrows/util/contexts"
	"github.com/teq-quocbang/arrows/util/myerror"
)

func (u *UseCase) Create(ctx context.Context, req *payload.CreateProductRequest) (*presenter.ProductResponseWrapper, error) {
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrProductInvalidParam(err.Error())
	}

	userPrinciple := contexts.GetUserPrincipleByContext(ctx)
	price, err := decimal.NewFromString(req.Price)
	if err != nil {
		return nil, myerror.ErrProductInvalidParam(err.Error())
	}
	producerID, err := uuid.Parse(req.ProducerID)
	if err != nil {
		return nil, myerror.ErrProductInvalidParam(err.Error())
	}

	product := &model.Product{
		Name:        req.Name,
		ProductType: req.ProductType,
		ProducerID:  producerID,
		Price:       price,
		CreatedBy:   userPrinciple.User.ID,
		UpdatedBy:   userPrinciple.User.ID,
	}
	if err := u.Product.Create(ctx, product); err != nil {
		return nil, myerror.ErrProductCreate(err)
	}

	return &presenter.ProductResponseWrapper{
		Product: product,
	}, nil
}
