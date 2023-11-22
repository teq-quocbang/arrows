package storage

import (
	"context"

	"github.com/teq-quocbang/arrows/payload"
	"github.com/teq-quocbang/arrows/presenter"
)

type IUseCase interface {
	UpsertStorage(context.Context, *payload.UpsertStorageRequest) (*presenter.StorageResponseWrapper, error)
	GetList(context.Context, *payload.GetStorageByLocatRequest) (*presenter.ListStorageResponseWrapper, error)
}
