package comment

import (
	"github.com/labstack/echo/v4"

	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"

	"github.com/teq-quocbang/arrows/payload"
)

// Create
// @Summary Create reply comment
// @Description create a reply comment
// @Tags Comment
// @Accept  json
// @Produce json
// @Security no
// @Param req body payload.CreateCommentRequest
// @Success 200 {object} presenter.CommentResponseWrapper
// @Router /comment/reply [comment] .
func (r *Route) ReplyComment(c echo.Context) error {
	var (
		ctx = &teq.CustomEchoContext{Context: c}
		req = &payload.ReplyCommentRequest{}
	)

	if err := c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
	}

	err := r.UseCase.Comment.ReplyComment(ctx, req)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, nil)
}
