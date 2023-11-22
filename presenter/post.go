package presenter

import "github.com/teq-quocbang/arrows/model"

type ReviewInfo struct {
	IsReacted    bool   `json:"is_reacted"`
	ReactedState string `json:"reacted_state"`
}

type PostResponseWrapper struct {
	Post   *model.Post `json:"Post"`
	Review ReviewInfo  `json:"review"`
}

type ListPostResponseWrapper struct {
	Post []model.Post `json:"Post"`
	Meta interface{}  `json:"meta"`
}
