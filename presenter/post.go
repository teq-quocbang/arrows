package presenter

import "github.com/teq-quocbang/store/model"

type PostResponseWrapper struct {
	Post *model.Post `json:"Post"`
}

type ListPostResponseWrapper struct {
	Post []model.Post `json:"Post"`
	Meta interface{}  `json:"meta"`
}
