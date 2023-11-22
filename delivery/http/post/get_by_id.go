package post

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"

	"github.com/teq-quocbang/store/presenter"
)

// GetByID post by id
// @Summary Get an post
// @Description Get post by id
// @Tags Example
// @Accept json
// @Produce json
// @Security AuthToken
// @Param id path int true "id"
// @Success 200 {object} presenter.PostResponseWrapper
// @Router /post/{id} [get] .
func (r *Route) GetByID(c echo.Context) error {
	var (
		ctx   = &teq.CustomEchoContext{Context: c}
		idStr = c.Param("id")
		resp  *presenter.PostResponseWrapper
	)

	postID, err := uuid.Parse(idStr)
	if err != nil {
		return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
	}

	resp, err = r.UseCase.Post.GetByID(ctx, postID)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}
