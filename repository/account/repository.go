package account

import (
	"context"

	"github.com/google/uuid"
	"github.com/teq-quocbang/arrows/model"
)

type Repository interface {
	CreateAccount(context.Context, *model.Account) (ID uuid.UUID, err error)
	GetAccountByUsername(ctx context.Context, username string) (*model.Account, error)
	GetAccountByConstraint(context.Context, *model.Account) (*model.Account, error)
	GetList(context.Context) ([]model.Account, error)
	GetByID(context.Context, uuid.UUID) (model.Account, error)
}
