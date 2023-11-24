package presenter

import (
	"time"

	"github.com/google/uuid"
	"github.com/teq-quocbang/arrows/model"
)

type CommentDetails struct {
	ID            uuid.UUID        `json:"id"`
	Content       string           `json:"content"`
	CreatedAt     time.Time        `json:"created_at"`
	CreatedBy     string           `json:"created_by"`
	ChildComments []CommentDetails `json:"child_comments"`
}

type ReviewInfo struct {
	IsReacted    bool             `json:"is_reacted"`
	ReactedState string           `json:"reacted_state"`
	Comments     []CommentDetails `json:"comment"`
}

type PostResponseWrapper struct {
	Post   *model.Post `json:"Post"`
	Review ReviewInfo  `json:"review"`
}

type ListPostResponseWrapper struct {
	Posts []PostResponseWrapper `json:"Post"`
	Meta  interface{}           `json:"meta"`
}
