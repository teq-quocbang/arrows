package post

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/teq-quocbang/arrows/payload"
	"github.com/teq-quocbang/arrows/presenter"
	"github.com/teq-quocbang/arrows/util/contexts"
	"github.com/teq-quocbang/arrows/util/myerror"
)

func (u *UseCase) GetList(ctx context.Context, req *payload.GetListPostRequest) (*presenter.ListPostResponseWrapper, error) {
	req.Format()
	var (
		userPrinciple = contexts.GetUserPrincipleByContext(ctx)
		order         = make([]string, 0)
	)

	if req.OrderBy != "" {
		order = append(order, fmt.Sprintf("%s %s", req.OrderBy, req.SortBy))
	}

	posts, total, err := u.Post.GetList(ctx, userPrinciple.User.ID, order, req.Paginator)
	if err != nil {
		return nil, myerror.ErrPostGet(err)
	}

	resp := &presenter.ListPostResponseWrapper{
		Posts: make([]presenter.PostResponseWrapper, len(posts)),
		Meta: map[string]interface{}{
			"page":  req.Paginator.Page,
			"limit": req.Paginator.Limit,
			"total": total,
		},
	}
	for idx, post := range posts {
		emoji, isReacted := post.ReactedThePost(userPrinciple.User.ID)
		comments, err := u.parseComment(ctx, post.Information.CommentIDs)
		if err != nil {
			return nil, err
		}
		resp.Posts[idx] = presenter.PostResponseWrapper{
			Post: &post,
			Review: presenter.ReviewInfo{
				IsReacted:    isReacted,
				ReactedState: string(emoji),
				Comments:     comments,
			},
		}
	}

	return resp, nil
}

func (u *UseCase) parseComment(ctx context.Context, commentIDs []uuid.UUID) ([]presenter.CommentDetails, error) {
	result := make([]presenter.CommentDetails, len(commentIDs))

	comments, err := u.Comment.GetByIDs(ctx, commentIDs)
	if err != nil {
		return nil, myerror.ErrCommentGet(err)
	}

	for i, comment := range comments {
		account, err := u.Account.GetByID(ctx, comment.CreatedBy)
		if err != nil {
			return nil, myerror.ErrAccountGet(err)
		}

		childCommentDetails, err := u.parseComment(ctx, comment.Information.ChildCommentIDs)
		if err != nil {
			return nil, err
		}

		result[i] = presenter.CommentDetails{
			ID:            comment.ID,
			Content:       comment.Contents,
			CreatedAt:     comment.CreatedAt,
			CreatedBy:     account.Username,
			ChildComments: childCommentDetails,
		}
	}

	return result, nil
}
