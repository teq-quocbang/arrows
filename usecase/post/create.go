package post

import (
	"context"

	"github.com/teq-quocbang/arrows/model"
	"github.com/teq-quocbang/arrows/payload"
	"github.com/teq-quocbang/arrows/presenter"
	"github.com/teq-quocbang/arrows/proto"
	"github.com/teq-quocbang/arrows/util/contexts"
	"github.com/teq-quocbang/arrows/util/myerror"
)

func (u *UseCase) Create(ctx context.Context, req *payload.CreatePostRequest) (*presenter.PostResponseWrapper, error) {
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrPostInvalidParam(err.Error())
	}

	userPrinciple := contexts.GetUserPrincipleByContext(ctx)

	post := &model.Post{
		Content:     req.Content,
		PrivacyMode: proto.Privacy(req.PrivacyMode),
		CreatedBy:   userPrinciple.User.ID,
	}
	if err := u.Post.Create(ctx, post); err != nil {
		return nil, myerror.ErrPostCreate(err)
	}

	return &presenter.PostResponseWrapper{
		Post: post,
	}, nil
}
