package post

import (
	"context"

	"github.com/google/uuid"
	"github.com/teq-quocbang/store/codetype"
	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/proto"
)

type Repository interface {
	Create(context.Context, *model.Post) error
	Update(ctx context.Context, postID uuid.UUID, content string, privacyMode proto.Privacy) error
	GetByID(context.Context, uuid.UUID) (model.Post, error)
	GetList(
		ctx context.Context,
		accountID uuid.UUID,
		order []string,
		paginator codetype.Paginator,
	) ([]model.Post, int64, error)
	UpsertEmoji(context.Context, uuid.UUID, *model.React) error
	Delete(context.Context, uuid.UUID) error
}
