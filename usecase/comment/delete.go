package comment

import (
	"context"
	"reflect"

	"github.com/google/uuid"
	"github.com/teq-quocbang/arrows/model"
	"github.com/teq-quocbang/arrows/proto"
	"github.com/teq-quocbang/arrows/util/contexts"
	"github.com/teq-quocbang/arrows/util/myerror"
)

func (u *UseCase) validateDelete(ctx context.Context, id uuid.UUID) (model.Comment, error) {
	if reflect.DeepEqual(id, uuid.UUID{}) {
		return model.Comment{}, myerror.ErrCommentInvalidParam("missing id")
	}

	userPrinciple := contexts.GetUserPrincipleByContext(ctx)

	// get comment
	comment, err := u.Comment.GetByID(ctx, id)
	if err != nil {
		return model.Comment{}, myerror.ErrCommentGet(err)
	}

	// get post
	post, err := u.Post.GetByID(ctx, comment.PostID)
	if err != nil {
		return model.Comment{}, myerror.ErrPostGet(err)
	}

	// check privacy
	if post.PrivacyMode != proto.Privacy_Public {
		return model.Comment{}, myerror.ErrPostForbidden("access denied")
	}

	// check whether is owner
	if comment.CreatedBy != userPrinciple.User.ID {
		return model.Comment{}, myerror.ErrCommentForbidden("not owner")
	}

	return comment, nil
}

func (u *UseCase) Delete(ctx context.Context, id uuid.UUID) error {
	comment, err := u.validateDelete(ctx, id)
	if err != nil {
		return err
	}

	// parent comment
	if comment.IsParent {
		if err := u.Comment.DeleteParent(ctx, comment); err != nil {
			return myerror.ErrCommentDelete(err)
		}
	} else { // child comment
		if err := u.Comment.DeleteChild(ctx, id, comment.Information.ParentID); err != nil {
			return myerror.ErrCommentDelete(err)
		}
	}
	return nil
}
