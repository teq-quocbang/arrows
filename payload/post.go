package payload

import "github.com/go-playground/validator/v10"

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
