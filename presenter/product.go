package presenter

import "github.com/teq-quocbang/arrows/model"

type ProductResponseWrapper struct {
	Product *model.Product `json:"product"`
}

type ListProductResponseWrapper struct {
	Product []model.Product `json:"product" yaml:"product"`
	Meta    interface{}     `json:"meta"`
}
