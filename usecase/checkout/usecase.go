package checkout

import (
	"github.com/teq-quocbang/arrows/config"
	"github.com/teq-quocbang/arrows/repository"
	"github.com/teq-quocbang/arrows/repository/checkout"
	"github.com/teq-quocbang/arrows/repository/product"
	"github.com/teq-quocbang/arrows/repository/storage"
	mySES "github.com/teq-quocbang/arrows/util/ses"
)

type UseCase struct {
	Product  product.Repository
	Checkout checkout.Repository
	Storage  storage.Repository

	SES mySES.ISES

	Config *config.Config
}

func New(repo *repository.Repository, ses mySES.ISES) IUseCase {
	return &UseCase{
		Product:  repo.Product,
		Checkout: repo.Checkout,
		Storage:  repo.Storage,
		SES:      ses,
		Config:   config.GetConfig(),
	}
}
