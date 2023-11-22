package grpc

import (
	"context"

	"github.com/teq-quocbang/arrows/model"
)

type IUseCase interface {
	GetByID(ctx context.Context, req *GetByIDRequest) (*model.Example, error)
}
