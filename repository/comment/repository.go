package comment

import (
	"context"

	"github.com/google/uuid"
	"github.com/teq-quocbang/arrows/model"
)

type Repository interface {
	CreateCommentInPost(context.Context, *model.Comment, *model.Post) error
	CreateInParentComment(ctx context.Context, cChild *model.Comment, cParent *model.Comment) error
	GetByID(context.Context, uuid.UUID) (model.Comment, error)
	Update(ctx context.Context, commentID uuid.UUID, contents string) error
	GetByIDs(ctx context.Context, IDs []uuid.UUID) ([]model.Comment, error)
	UpsertEmoji(context.Context, uuid.UUID, *model.React) error
	DeleteChild(ctx context.Context, cChildID uuid.UUID, cParentID uuid.UUID) error
	DeleteParent(ctx context.Context, c model.Comment) error
}
