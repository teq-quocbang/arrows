package presenter

import "github.com/teq-quocbang/arrows/model"

type ExampleResponseWrapper struct {
	Example *model.Example `json:"example"`
}

type ListExampleResponseWrapper struct {
	Examples []model.Example `json:"examples"`
	Meta     interface{}     `json:"meta"`
}
