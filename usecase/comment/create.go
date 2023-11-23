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

func (u *UseCase) ReplyComment(ctx context.Context, req *payload.ReplyCommentRequest) error {
	if err := req.Validate(); err != nil {
		return myerror.ErrCommentInvalidParam(err.Error())
	}

	userPrinciple := contexts.GetUserPrincipleByContext(ctx)

	pCommentID, err := uuid.Parse(req.ParentCommentID)
	if err != nil {
		return myerror.ErrCommentInvalidParam(err.Error())
	}

	// get parent
	pComment, err := u.Comment.GetByID(ctx, pCommentID)
	if err != nil {
		return myerror.ErrCommentGet(err)
	}

	// get post
	post, err := u.Post.GetByID(ctx, pComment.PostID)
	if err != nil {
		return myerror.ErrPostGet(err)
	}

	// check post privacy mode
	if post.PrivacyMode != proto.Privacy_Public {
		if userPrinciple.User.ID != post.CreatedBy {
			return myerror.ErrPostForbidden("access denied")
		}
	}

	// check whether is parent comment
	if !pComment.IsParent {
		return myerror.ErrCommentCannotReply("is not parent comment")
	}

	// create child comment
	cComment := &model.Comment{
		Contents: req.Content,
		IsParent: false,
		Information: model.CommentInfo{
			ParentID: pCommentID,
		},
		CreatedBy: userPrinciple.User.ID,
		PostID:    pComment.PostID,
	}
	if err := u.Comment.CreateInParentComment(ctx, cComment, &pComment); err != nil {
		return myerror.ErrCommentCreate(err)
	}

	return nil
}
