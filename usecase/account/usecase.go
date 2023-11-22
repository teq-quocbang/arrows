package account

import (
	"github.com/teq-quocbang/arrows/config"
	"github.com/teq-quocbang/arrows/repository"
	"github.com/teq-quocbang/arrows/repository/account"
	mySES "github.com/teq-quocbang/arrows/util/ses"
)

type UseCase struct {
	Account account.Repository

	SES mySES.ISES

	Config *config.Config
}

func New(repo *repository.Repository, ses mySES.ISES) IUseCase {
	return &UseCase{
		Account: repo.Account,
		SES:     ses,
		Config:  config.GetConfig(),
	}
}
