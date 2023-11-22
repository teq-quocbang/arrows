package statistics

import (
	"github.com/teq-quocbang/arrows/config"
	"github.com/teq-quocbang/arrows/repository"
	"github.com/teq-quocbang/arrows/repository/checkout"
	"github.com/teq-quocbang/arrows/repository/producer"
	"github.com/teq-quocbang/arrows/repository/product"
	"github.com/teq-quocbang/arrows/repository/storage"
	mySES "github.com/teq-quocbang/arrows/util/ses"
)

type UseCase struct {
	Producer producer.Repository
	Storage  storage.Repository
	Checkout checkout.Repository
	Product  product.Repository

	SES mySES.ISES

	Config *config.Config
}

func New(repo *repository.Repository, ses mySES.ISES) IUseCase {
	return &UseCase{
		Producer: repo.Producer,
		SES:      ses,
		Storage:  repo.Storage,
		Checkout: repo.Checkout,
		Product:  repo.Product,
		Config:   config.GetConfig(),
	}
}
