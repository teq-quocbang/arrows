package grpc

import (
	"github.com/teq-quocbang/arrows/config"
	"github.com/teq-quocbang/arrows/repository"
	"github.com/teq-quocbang/arrows/repository/example"
)

type UseCase struct {
	ExampleRepo example.Repository

	Config *config.Config
}

func New(repo *repository.Repository) IUseCase {
	return &UseCase{
		ExampleRepo: repo.Example,
		Config:      config.GetConfig(),
	}
}
