package product

import (
	"github.com/teq-quocbang/arrows/config"
	"github.com/teq-quocbang/arrows/repository"
	"github.com/teq-quocbang/arrows/repository/product"
	mySES "github.com/teq-quocbang/arrows/util/ses"
)

type UseCase struct {
	Product product.Repository

	SES mySES.ISES

	Config *config.Config
}

func New(repo *repository.Repository, ses mySES.ISES) IUseCase {
	return &UseCase{
		Product: repo.Product,
		SES:     ses,
		Config:  config.GetConfig(),
	}
}
