package account

import (
	"context"

	"github.com/teq-quocbang/arrows/payload"
	"github.com/teq-quocbang/arrows/presenter"
)

type IUseCase interface {
	SignUp(context.Context, *payload.SignUpRequest) (*presenter.AccountResponseWrapper, error)
	Login(context.Context, *payload.LoginRequest) (*presenter.AccountLoginResponseWrapper, error)
}
