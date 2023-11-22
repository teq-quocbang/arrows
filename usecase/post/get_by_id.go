package post

import (
	"context"
	"reflect"

	"github.com/google/uuid"

	"github.com/teq-quocbang/arrows/presenter"
	"github.com/teq-quocbang/arrows/proto"
	"github.com/teq-quocbang/arrows/util/contexts"
	"github.com/teq-quocbang/arrows/util/myerror"
)

func (u *UseCase) GetByID(ctx context.Context, id uuid.UUID) (*presenter.PostResponseWrapper, error) {
	if reflect.DeepEqual(id, uuid.UUID{}) {
		return nil, myerror.ErrPostInvalidParam("missing id")
	}

	userPrinciple := contexts.GetUserPrincipleByContext(ctx)

	post, err := u.Post.GetByID(ctx, id)
	if err != nil {
		return nil, myerror.ErrPostGet(err)
	}

	// check whether this post is private
	if post.PrivacyMode != proto.Privacy_Public {
		if post.CreatedBy != userPrinciple.User.ID {
			return nil, myerror.ErrPostForbidden("access denied")
		}
	}

	return &presenter.PostResponseWrapper{
		Post: &post,
	}, nil
}
