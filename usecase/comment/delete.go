package comment

import (
	"context"
	"reflect"

	"github.com/google/uuid"
	"github.com/teq-quocbang/arrows/proto"
	"github.com/teq-quocbang/arrows/util/contexts"
	"github.com/teq-quocbang/arrows/util/myerror"
)

func (u *UseCase) Delete(ctx context.Context, id uuid.UUID) error {
	if reflect.DeepEqual(id, uuid.UUID{}) {
		return myerror.ErrCommentInvalidParam("missing id")
	}

	userPrinciple := contexts.GetUserPrincipleByContext(ctx)

	// get comment
	comment, err := u.Comment.GetByID(ctx, id)
	if err != nil {
		return myerror.ErrCommentGet(err)
	}

	// get post
	post, err := u.Post.GetByID(ctx, comment.PostID)
	if err != nil {
		return myerror.ErrPostGet(err)
	}

	// check privacy
	if post.PrivacyMode != proto.Privacy_Public {
		return myerror.ErrPostForbidden("access denied")
	}

	// check whether is owner
	if comment.CreatedBy != userPrinciple.User.ID {
		return myerror.ErrCommentForbidden("not owner")
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
