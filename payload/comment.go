package payload

import "github.com/go-playground/validator/v10"

type CreateCommentRequest struct {
	PostID  string `json:"post_id" validate:"required"`
	Content string `json:"content" validate:"required"`
}

func (c *CreateCommentRequest) Validate() error {
	var validate = validator.New()
	return validate.Struct(c)
}

type ReplyCommentRequest struct {
	ParentCommentID string `json:"parent_comment_id" validate:"required"`
	Content         string `json:"content" validate:"required"`
}

func (r *ReplyCommentRequest) Validate() error {
	var validate = validator.New()
	return validate.Struct(r)
}
