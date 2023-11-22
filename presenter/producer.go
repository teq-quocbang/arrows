package presenter

import "github.com/teq-quocbang/arrows/model"

type ProducerResponseWrapper struct {
	Producer *model.Producer `json:"producer"`
}

type ListProducerResponseWrapper struct {
	Producer []model.Producer `json:"producer"`
	Meta     interface{}      `json:"meta"`
}
