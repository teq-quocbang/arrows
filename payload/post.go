package payload

import (
	"github.com/go-playground/validator/v10"
	"github.com/teq-quocbang/arrows/codetype"
)

type CreatePostRequest struct {
	Content     string `json:"content" validate:"required"`
	PrivacyMode int32  `json:"privacy_mode" validate:"required"`
}

func (p *CreatePostRequest) Validate() error {
	var validate = validator.New()
	return validate.Struct(p)
}

type UpsertEmojiRequest struct {
	PostID string `json:"post_id" validate:"required"`
	Emoji  string `json:"emoji" validate:"required"`
}

func (p *UpsertEmojiRequest) Validate() error {
	var validate = validator.New()
	return validate.Struct(p)
}

type UpdatePostRequest struct {
	PostID      string `json:"post_id" validate:"required"`
	Content     string `json:"content"`
	PrivacyMode int32  `json:"privacy_mode"`
}

func (p *UpdatePostRequest) IsNoUpdate() bool {
	return p.Content == "" && p.PrivacyMode == 0
}

func (p *UpdatePostRequest) Validate() error {
	var validate = validator.New()
	return validate.Struct(p)
}

type GetListPostRequest struct {
	codetype.Paginator
	SortBy  codetype.SortType
	OrderBy string `json:"order_by,omitempty" query:"order_by"`
}
