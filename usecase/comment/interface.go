package comment

import (
	"context"

	"github.com/teq-quocbang/arrows/payload"
)

type IUseCase interface {
	Create(context.Context, *payload.CreateCommentRequest) error
	ReplyComment(context.Context, *payload.ReplyCommentRequest) error
}
