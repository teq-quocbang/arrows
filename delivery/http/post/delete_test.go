package post

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/teq-quocbang/arrows/delivery/http/auth"
	"github.com/teq-quocbang/arrows/fixture/database"
	"github.com/teq-quocbang/arrows/repository"
	"github.com/teq-quocbang/arrows/usecase"
	"github.com/teq-quocbang/arrows/util/token"
)

func TestDelete(t *testing.T) {
	assertion := assert.New(t)
	db := database.InitDatabase()
	defer db.TruncateTables()

	repo := repository.New(db.GetClient)
	r := Route{
		UseCase: usecase.New(repo, nil),
	}

	accountID, postID, err := setUpDependencyData(db)
	assertion.NoError(err)

	userPrinciple := &token.JWTClaimCustom{
		SessionID: uuid.New(),
		User: token.UserInfo{
			Username: gofakeit.Name(),
			ID:       accountID,
			Email:    gofakeit.Email(),
		},
	}

	// good case
	{
		// Arrange
		resp, ctx := setupDelete(postID.String())
		ctx.Set(string(auth.UserPrincipleKey), userPrinciple)

		// Act
		err := r.Delete(ctx)

		// Assert
		assertion.NoError(err)
		assertion.Equal(200, resp.Code)
	}

	// bad case
	{
		// Arrange
		resp, ctx := setupDelete("")
		ctx.Set(string(auth.UserPrincipleKey), userPrinciple)

		// Act
		r.Delete(ctx)

		// Assert
		assertion.Equal(400, resp.Code)
	}
}

func setupDelete(id string) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/api/post/:id", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(id)

	return rec, c
}
