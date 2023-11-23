package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/teq-quocbang/arrows/repository/account"
	"github.com/teq-quocbang/arrows/repository/checkout"
	"github.com/teq-quocbang/arrows/repository/comment"
	"github.com/teq-quocbang/arrows/repository/example"
	"github.com/teq-quocbang/arrows/repository/post"
	"github.com/teq-quocbang/arrows/repository/producer"
	"github.com/teq-quocbang/arrows/repository/product"
	"github.com/teq-quocbang/arrows/repository/storage"
)

type Repository struct {
	Account  account.Repository
	Example  example.Repository
	Product  product.Repository
	Producer producer.Repository
	Storage  storage.Repository
	Checkout checkout.Repository
	Post     post.Repository
	Comment  comment.Repository
}

func New(getClient func(ctx context.Context) *gorm.DB) *Repository {
	return &Repository{
		Account:  account.NewAccountPG(getClient),
		Example:  example.NewPG(getClient),
		Product:  product.NewPG(getClient),
		Producer: producer.NewPG(getClient),
		Storage:  storage.NewPG(getClient),
		Checkout: checkout.NewPG(getClient),
		Post:     post.NewPG(getClient),
		Comment:  comment.NewPG(getClient),
	}
}
