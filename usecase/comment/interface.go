package comment

import (
	"context"

	"github.com/google/uuid"

	"github.com/teq-quocbang/arrows/payload"
	"github.com/teq-quocbang/arrows/presenter"
)

type IUseCase interface {
	Create(context.Context, *payload.CreateCommentRequest) error
	ReplyComment(context.Context, *payload.ReplyCommentRequest) error
	Update(context.Context, *payload.UpdateCommentRequest) (*presenter.CommentResponseWrapper, error)
	Delete(context.Context, uuid.UUID) error
}
