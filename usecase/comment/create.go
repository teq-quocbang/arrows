package comment

import (
	"context"

	"github.com/google/uuid"
	"github.com/teq-quocbang/arrows/model"
	"github.com/teq-quocbang/arrows/payload"
	"github.com/teq-quocbang/arrows/proto"
	"github.com/teq-quocbang/arrows/util/contexts"
	"github.com/teq-quocbang/arrows/util/myerror"
)

func (u *UseCase) Create(ctx context.Context, req *payload.CreateCommentRequest) error {
	if err := req.Validate(); err != nil {
		return myerror.ErrCommentInvalidParam(err.Error())
	}

	postID, err := uuid.Parse(req.PostID)
	if err != nil {
		return myerror.ErrCommentInvalidParam(err.Error())
	}

	post, err := u.Post.GetByID(ctx, postID)
	if err != nil {
		return myerror.ErrPostGet(err)
	}

	userPrinciple := contexts.GetUserPrincipleByContext(ctx)
	// check privacy
	if post.PrivacyMode != proto.Privacy_Public {
		if post.CreatedBy != userPrinciple.User.ID {
			return myerror.ErrCommentForbidden("access denied")
		}
	}

	// create comment
	comment := &model.Comment{
		Contents:  req.Content,
		IsParent:  true,
		PostID:    postID,
		CreatedBy: userPrinciple.User.ID,
	}
	if err := u.Comment.CreateCommentInPost(ctx, comment, &post); err != nil {
		return myerror.ErrCommentCreate(err)
	}

	return nil
}
