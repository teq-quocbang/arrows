package statistics

import (
	"context"

	"github.com/teq-quocbang/arrows/payload"
	"github.com/teq-quocbang/arrows/presenter"
)

type IUseCase interface {
	GetProductSoldChart(context.Context, *payload.GetChartRequest) (*presenter.ListStatisticsSoldProductChartResponseWrapper, error)
	GetProductGrowthChart(context.Context, *payload.GetChartRequest) (*presenter.ListStatisticsSoldProductChartResponseWrapper, error)
}
