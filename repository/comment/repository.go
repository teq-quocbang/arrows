package comment

import (
	"context"

	"github.com/google/uuid"
	"github.com/teq-quocbang/arrows/model"
)

type Repository interface {
	CreateCommentInPost(context.Context, *model.Comment, *model.Post) error
	CreateInParentComment(ctx context.Context, parentID uuid.UUID, comment model.Comment)
	Update(ctx context.Context, commentID uuid.UUID, contents string) error
	UpsertEmoji(context.Context, uuid.UUID, *model.React) error
	Delete(context.Context, uuid.UUID) error
}
