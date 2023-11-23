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
