package comment

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
	group.POST("/reply", r.ReplyComment)
	group.PUT("", r.Update)
	group.DELETE("/:id", r.Delete)
	group.PATCH("/react", r.ReactEmoji)
}
