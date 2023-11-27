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

func (u *UseCase) validateReplyComment(ctx context.Context, req *payload.ReplyCommentRequest, userID uuid.UUID) (*model.Comment, error) {
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrCommentInvalidParam(err.Error())
	}

	pCommentID, err := uuid.Parse(req.ParentCommentID)
	if err != nil {
		return nil, myerror.ErrCommentInvalidParam(err.Error())
	}

	// get parent
	pComment, err := u.Comment.GetByID(ctx, pCommentID)
	if err != nil {
		return nil, myerror.ErrCommentGet(err)
	}

	// get post
	post, err := u.Post.GetByID(ctx, pComment.PostID)
	if err != nil {
		return nil, myerror.ErrPostGet(err)
	}

	// check post privacy mode
	if post.PrivacyMode != proto.Privacy_Public {
		if userID != post.CreatedBy {
			return nil, myerror.ErrPostForbidden("access denied")
		}
	}

	// check whether is parent comment
	if !pComment.IsParent {
		return nil, myerror.ErrCommentCannotReply("is not parent comment")
	}

	return &pComment, nil
}

func (u *UseCase) ReplyComment(ctx context.Context, req *payload.ReplyCommentRequest) error {
	userPrinciple := contexts.GetUserPrincipleByContext(ctx)
	pComment, err := u.validateReplyComment(ctx, req, userPrinciple.User.ID)
	if err != nil {
		return err
	}

	// create child comment
	cComment := &model.Comment{
		Contents: req.Content,
		IsParent: false,
		Information: model.CommentInfo{
			Parent: model.ParentInfo{
				ParentID: pComment.ID.String(),
			},
		},
		CreatedBy: userPrinciple.User.ID,
		PostID:    pComment.PostID,
	}
	if err := u.Comment.CreateInParentComment(ctx, cComment, pComment); err != nil {
		return myerror.ErrCommentCreate(err)
	}

	return nil
}
