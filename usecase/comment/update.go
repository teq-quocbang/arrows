package comment

import (
	"context"
	"slices"

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

func (u *UseCase) ReactEmoji(ctx context.Context, req *payload.ReactEmojiRequest) (*presenter.CommentResponseWrapper, error) {
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrCommentInvalidParam(err.Error())
	}

	userPrinciple := contexts.GetUserPrincipleByContext(ctx)

	commentID, err := uuid.Parse(req.CommentID)
	if err != nil {
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
		if userPrinciple.User.ID != post.CreatedBy {
			return nil, myerror.ErrPostForbidden("access denied")
		}
	}

	reactUpsert := model.React{}
	if comment.Information.Reacts != nil {
		// remove old react
		for emoji, userIDs := range comment.Information.Reacts {
			// find old react
			if slices.Contains[[]uuid.UUID](userIDs, userPrinciple.User.ID) {
				// check whether is same emoji
				if emoji == model.Emoji(req.Emoji) {
					// clear nil emoji
					comment.Information.Reacts.ClearNilReact()
					return &presenter.CommentResponseWrapper{
						Comment: &comment,
						Review: presenter.ReviewInfo{
							IsReacted:    true,
							ReactedState: req.Emoji,
						},
					}, nil
				}
				// remove old react
				reactUpsert[emoji] = slices.DeleteFunc[[]uuid.UUID](userIDs, func(u uuid.UUID) bool {
					return userPrinciple.User.ID == u
				})

				// update comment
				comment.Information.Reacts[emoji] = reactUpsert[emoji]
			}
		}

		//let new react append old userID
		comment.Information.Reacts[model.Emoji(req.Emoji)] = append(comment.Information.Reacts[model.Emoji(req.Emoji)], userPrinciple.User.ID)
		reactUpsert[model.Emoji(req.Emoji)] = comment.Information.Reacts[model.Emoji(req.Emoji)]
	} else {
		reactUpsert[model.Emoji(req.Emoji)] = []uuid.UUID{userPrinciple.User.ID}
		comment.Information.Reacts = reactUpsert // update react
	}

	// react emoji
	if err := u.Comment.UpsertEmoji(ctx, commentID, &reactUpsert); err != nil {
		return nil, myerror.ErrCommentUpsert(err)
	}

	comment.Information.Reacts.ClearNilReact()
	return &presenter.CommentResponseWrapper{
		Comment: &comment,
		Review: presenter.ReviewInfo{
			IsReacted:    true,
			ReactedState: req.Emoji,
		},
	}, nil
}
