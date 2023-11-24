package presenter

import "github.com/teq-quocbang/arrows/model"

type CommentResponseWrapper struct {
	Comment *model.Comment `json:"Comment"`
	Review  ReviewInfo     `json:"review"`
}

type ListCommentResponseWrapper struct {
	Comment []model.Comment `json:"Comment"`
	Meta    interface{}     `json:"meta"`
}
