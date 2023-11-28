package post

import (
	"github.com/teq-quocbang/arrows/config"
	"github.com/teq-quocbang/arrows/repository"
	"github.com/teq-quocbang/arrows/repository/account"
	"github.com/teq-quocbang/arrows/repository/comment"
	"github.com/teq-quocbang/arrows/repository/post"
	mySES "github.com/teq-quocbang/arrows/util/ses"
)

type UseCase struct {
	Post    post.Repository
	Comment comment.Repository
	Account account.Repository
	SES     mySES.ISES

	Config *config.Config
}

func New(repo *repository.Repository, ses mySES.ISES) IUseCase {
	return &UseCase{
		Post:    repo.Post,
		Comment: repo.Comment,
		Account: repo.Account,
		SES:     ses,
		Config:  config.GetConfig(),
	}
}
