package storage

import (
	"context"

	"github.com/teq-quocbang/arrows/payload"
	"github.com/teq-quocbang/arrows/presenter"
	"github.com/teq-quocbang/arrows/util/myerror"
)

func (u *UseCase) GetList(ctx context.Context, req *payload.GetStorageByLocatRequest) (*presenter.ListStorageResponseWrapper, error) {
	if err := req.Validate(); err != nil {
		return &presenter.ListStorageResponseWrapper{}, nil
	}

	storages, err := u.Storage.GetListStorageByLocat(ctx, req.Locat)
	if err != nil {
		return nil, myerror.ErrStorageGet(err)
	}

	return &presenter.ListStorageResponseWrapper{
		Storage: storages,
	}, nil
}
