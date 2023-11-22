package post

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
	"github.com/teq-quocbang/arrows/util/replace"
)

func (u *UseCase) UpsertEmoji(ctx context.Context, req *payload.UpsertEmojiRequest) (*presenter.PostResponseWrapper, error) {
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrPostInvalidParam(err.Error())
	}

	userPrinciple := contexts.GetUserPrincipleByContext(ctx)

	postID, err := uuid.Parse(req.PostID)
	if err != nil {
		return nil, myerror.ErrPostInvalidParam(err.Error())
	}

	// get post
	post, err := u.Post.GetByID(ctx, postID)
	if err != nil {
		return nil, myerror.ErrPostGet(err)
	}

	// check permission
	if post.PrivacyMode != proto.Privacy_Public {
		if post.CreatedBy != userPrinciple.User.ID {
			return nil, myerror.ErrPostForbidden("access denied")
		}
	}

	reactUpsert := model.React{}
	if post.Information.Reacts != nil {
		// remove old react
		for emoji, userIDs := range post.Information.Reacts {
			// find old react
			if slices.Contains[[]uuid.UUID](userIDs, userPrinciple.User.ID) {
				// check whether is same emoji
				if emoji == model.Emoji(req.Emoji) {
					// clear nil emoji
					post.Information.Reacts.ClearNilReact()
					return &presenter.PostResponseWrapper{
						Post: &post,
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

				// update post
				post.Information.Reacts[emoji] = reactUpsert[emoji]
			}
		}

		//let new react append old userID
		post.Information.Reacts[model.Emoji(req.Emoji)] = append(post.Information.Reacts[model.Emoji(req.Emoji)], userPrinciple.User.ID)
		reactUpsert[model.Emoji(req.Emoji)] = post.Information.Reacts[model.Emoji(req.Emoji)]
	} else {
		reactUpsert[model.Emoji(req.Emoji)] = []uuid.UUID{userPrinciple.User.ID}
		post.Information.Reacts = reactUpsert // update react
	}

	// upsert emoji
	if err := u.Post.UpsertEmoji(ctx, postID, &reactUpsert); err != nil {
		return nil, myerror.ErrPostUpdate(err)
	}

	// clear nil emoji
	post.Information.Reacts.ClearNilReact()
	return &presenter.PostResponseWrapper{
		Post: &post,
		Review: presenter.ReviewInfo{
			IsReacted:    true,
			ReactedState: req.Emoji,
		},
	}, nil
}

func (u *UseCase) Update(ctx context.Context, req *payload.UpdatePostRequest) (*presenter.PostResponseWrapper, error) {
	if req.IsNoUpdate() {
		return nil, myerror.ErrPostInvalidParam("no request update params")
	}
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrPostInvalidParam(err.Error())
	}

	userPrinciple := contexts.GetUserPrincipleByContext(ctx)
	postID, err := uuid.Parse(req.PostID)
	if err != nil {
		return nil, myerror.ErrPostInvalidParam(err.Error())
	}

	// get post
	post, err := u.Post.GetByID(ctx, postID)
	if err != nil {
		return nil, myerror.ErrPostGet(err)
	}

	// check permission
	if post.CreatedBy != userPrinciple.User.ID {
		return nil, myerror.ErrPostForbidden("not owner")
	}

	post, err = replace.Replace(model.Post{
		Content:     req.Content,
		PrivacyMode: proto.Privacy(req.PrivacyMode),
	}, post)
	if err != nil {
		return nil, myerror.ErrPostUpdate(err)
	}

	if err := u.Post.Update(ctx, postID, post.Content, post.PrivacyMode); err != nil {
		return nil, myerror.ErrPostUpdate(err)
	}

	emoji, ok := post.ReactedThePost(userPrinciple.User.ID)
	// clear nil emoji
	post.Information.Reacts.ClearNilReact()
	return &presenter.PostResponseWrapper{
		Post: &post,
		Review: presenter.ReviewInfo{
			IsReacted:    ok,
			ReactedState: string(emoji),
		},
	}, nil
}
