package producer

import (
	"github.com/teq-quocbang/arrows/config"
	"github.com/teq-quocbang/arrows/repository"
	"github.com/teq-quocbang/arrows/repository/producer"
	mySES "github.com/teq-quocbang/arrows/util/ses"
)

type UseCase struct {
	Producer producer.Repository

	SES mySES.ISES

	Config *config.Config
}

func New(repo *repository.Repository, ses mySES.ISES) IUseCase {
	return &UseCase{
		Producer: repo.Producer,
		SES:      ses,
		Config:   config.GetConfig(),
	}
}
