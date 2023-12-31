package post

import (
	"context"

	"github.com/google/uuid"

	"github.com/teq-quocbang/arrows/payload"
	"github.com/teq-quocbang/arrows/presenter"
)

type IUseCase interface {
	Create(context.Context, *payload.CreatePostRequest) (*presenter.PostResponseWrapper, error)
	GetByID(context.Context, uuid.UUID) (*presenter.PostResponseWrapper, error)
	UpsertEmoji(context.Context, *payload.UpsertEmojiRequest) (*presenter.PostResponseWrapper, error)
	Update(context.Context, *payload.UpdatePostRequest) (*presenter.PostResponseWrapper, error)
	Delete(context.Context, uuid.UUID) error
	GetList(context.Context, *payload.GetListPostRequest) (*presenter.ListPostResponseWrapper, error)
}
