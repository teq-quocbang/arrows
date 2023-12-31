package post

import (
	"github.com/labstack/echo/v4"
	"github.com/teq-quocbang/arrows/usecase"
)

type Route struct {
	UseCase *usecase.UseCase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{UseCase: useCase}

	group.POST("", r.Create)
	group.GET("/:id", r.GetByID)
	group.PATCH("", r.UpsertEmoji)
	group.PUT("", r.Update)
	group.DELETE("/:id", r.Delete)
	group.GET("", r.GetList)
}
