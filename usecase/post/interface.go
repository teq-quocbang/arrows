package post

import (
	"context"

	"github.com/google/uuid"

	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
)

type IUseCase interface {
	Create(context.Context, *payload.CreatePostRequest) (*presenter.PostResponseWrapper, error)
	GetByID(context.Context, uuid.UUID) (*presenter.PostResponseWrapper, error)
}
