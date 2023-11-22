package producer

import (
	"context"

	"github.com/teq-quocbang/arrows/model"
)

type Repository interface {
	Create(context.Context, *model.Producer) error
}
