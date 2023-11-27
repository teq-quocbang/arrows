package comment

import (
	"context"

	"github.com/google/uuid"

	"github.com/teq-quocbang/arrows/model"
	"github.com/teq-quocbang/arrows/payload"
	"github.com/teq-quocbang/arrows/presenter"
	"github.com/teq-quocbang/arrows/proto"
	"github.com/teq-quocbang/arrows/util/contexts"
	"github.com/teq-quocbang/arrows/util/myerror"
)

func (u *UseCase) validateUpdate(ctx context.Context, req *payload.UpdateCommentRequest, commentID uuid.UUID, userID uuid.UUID) (*model.Comment, error) {
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrCommentInvalidParam(err.Error())
	}

	// get comment
	comment, err := u.Comment.GetByID(ctx, commentID)
	if err != nil {
		return nil, myerror.ErrCommentGet(err)
	}

	// get post
	post, err := u.Post.GetByID(ctx, comment.PostID)
	if err != nil {
		return nil, myerror.ErrPostGet(err)
	}

	// check privacy
	if post.PrivacyMode != proto.Privacy_Public {
		if userID != post.CreatedBy {
			return nil, myerror.ErrPostForbidden("access denied")
		}
	}

	// check is owner
	if comment.CreatedBy != userID {
		return nil, myerror.ErrCommentForbidden("not owner")
	}

	return &comment, nil
}

func (u *UseCase) Update(ctx context.Context, req *payload.UpdateCommentRequest) (*presenter.CommentResponseWrapper, error) {
	userPrinciple := contexts.GetUserPrincipleByContext(ctx)
	commentID, err := uuid.Parse(req.CommentID)
	if err != nil {
		return nil, myerror.ErrCommentInvalidParam(err.Error())
	}

	comment, err := u.validateUpdate(ctx, req, commentID, userPrinciple.User.ID)
	if err != nil {
		return nil, err
	}

	// let update
	if err := u.Comment.Update(ctx, commentID, req.Content); err != nil {
		return nil, myerror.ErrCommentUpdate(err)
	}

	emoji, ok := comment.ReactedThePost(userPrinciple.User.ID)
	comment.Contents = req.Content
	return &presenter.CommentResponseWrapper{
		Comment: comment,
		Review: presenter.ReviewInfo{
			IsReacted:    ok,
			ReactedState: string(emoji),
		},
	}, nil
}
