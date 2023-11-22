package post

import (
	"github.com/teq-quocbang/arrows/config"
	"github.com/teq-quocbang/arrows/repository"
	"github.com/teq-quocbang/arrows/repository/post"
	mySES "github.com/teq-quocbang/arrows/util/ses"
)

type UseCase struct {
	Post post.Repository
	SES  mySES.ISES

	Config *config.Config
}

func New(repo *repository.Repository, ses mySES.ISES) IUseCase {
	return &UseCase{
		Post:   repo.Post,
		SES:    ses,
		Config: config.GetConfig(),
	}
}
