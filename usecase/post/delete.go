package post

import (
	"context"
	"reflect"

	"github.com/google/uuid"

	"github.com/teq-quocbang/arrows/util/contexts"
	"github.com/teq-quocbang/arrows/util/myerror"
)

func (u *UseCase) Delete(ctx context.Context, id uuid.UUID) error {
	if reflect.DeepEqual(id, uuid.UUID{}) {
		return myerror.ErrPostInvalidParam("missing id")
	}

	userPrinciple := contexts.GetUserPrincipleByContext(ctx)
	post, err := u.Post.GetByID(ctx, id)
	if err != nil {
		return myerror.ErrPostGet(err)
	}

	// check whether is owner
	if post.CreatedBy != userPrinciple.User.ID {
		return myerror.ErrPostForbidden("not owner")
	}

	if err := u.Post.Delete(ctx, id); err != nil {
		return myerror.ErrPostDelete(err)
	}

	return nil
}
