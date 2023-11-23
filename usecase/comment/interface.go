package comment

import (
	"context"

	"github.com/teq-quocbang/arrows/payload"
	"github.com/teq-quocbang/arrows/presenter"
)

type IUseCase interface {
	Create(context.Context, *payload.CreateCommentRequest) (*presenter.PostResponseWrapper, error)
	ReplyComment(context.Context, *payload.ReplyCommentRequest) error
}
