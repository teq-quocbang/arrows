package presenter

import "github.com/teq-quocbang/arrows/model"

type StorageResponseWrapper struct {
	Storage *model.Storage `json:"product"`
}

type ListStorageResponseWrapper struct {
	Storage []model.Storage `json:"product" yaml:"product"`
	Meta    interface{}     `json:"meta"`
}
