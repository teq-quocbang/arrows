package post

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/teq-quocbang/store/delivery/http/account"
	"github.com/teq-quocbang/store/delivery/http/auth"
	"github.com/teq-quocbang/store/fixture/database"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/proto"
	"github.com/teq-quocbang/store/repository"
	"github.com/teq-quocbang/store/usecase"
	"github.com/teq-quocbang/store/util/test"
	"github.com/teq-quocbang/store/util/token"
)

func TestGetByID(t *testing.T) {
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
		resp, ctx := setupGetByID(postID.String())
		ctx.Set(string(auth.UserPrincipleKey), userPrinciple)

		// Act
		err = r.GetByID(ctx)

		// Assert
		assertion.NoError(err)
		actual, err := test.UnmarshalBody[*presenter.PostResponseWrapper](resp.Body.Bytes())
		assertion.NoError(err)
		assertion.Equal(postID, actual.Post.ID)
	}

	// bad case
	{
		// Arrange
		resp, ctx := setupGetByID("")
		ctx.Set(string(auth.UserPrincipleKey), userPrinciple)

		// Act
		r.GetByID(ctx)

		// Assert
		assertion.Equal(400, resp.Code)
	}
}

func setupGetByID(id string) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/post/:id", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(id)

	return rec, c
}

func setUpDependencyData(db *database.Database) (uuid.UUID, uuid.UUID, error) {
	repo := repository.New(db.GetClient)
	rAccount := account.Route{
		UseCase: usecase.New(repo, nil),
	}
	rPost := Route{
		UseCase: usecase.New(repo, nil),
	}

	// create account
	resp, ctx := setUpTestSignUp(&payload.SignUpRequest{
		Username: gofakeit.Name(),
		Password: gofakeit.Password(true, true, true, false, false, 10),
		Email:    gofakeit.Email(),
	})
	err := rAccount.SignUp(ctx)
	if err != nil {
		return uuid.UUID{}, uuid.UUID{}, err
	}
	if resp.Code != 200 {
		return uuid.UUID{}, uuid.UUID{}, fmt.Errorf("failed to sign up, error: %v", err)
	}
	account, err := test.UnmarshalBody[*presenter.AccountResponseWrapper](resp.Body.Bytes())
	if err != nil {
		return uuid.UUID{}, uuid.UUID{}, fmt.Errorf("failed to unmarshal sign up body, error: %v", err)
	}

	// set userPrinciple
	userPrinciple := &token.JWTClaimCustom{
		SessionID: uuid.New(),
		User: token.UserInfo{
			ID: account.Account.ID,
		},
	}

	// create post
	req := &payload.CreatePostRequest{
		Content:     "abc",
		PrivacyMode: int32(proto.Privacy_Public.Number()),
	}
	resp, ctx = setupCreate(req)
	ctx.Set(string(auth.UserPrincipleKey), userPrinciple)
	err = rPost.Create(ctx)
	if err != nil {
		return uuid.UUID{}, uuid.UUID{}, fmt.Errorf("failed to create post, error: %v", err)
	}
	if resp.Code != 200 {
		return uuid.UUID{}, uuid.UUID{}, fmt.Errorf("failed to create post, error: %v", resp.Body)
	}
	post, err := test.UnmarshalBody[*presenter.PostResponseWrapper](resp.Body.Bytes())
	if err != nil {
		return uuid.UUID{}, uuid.UUID{}, fmt.Errorf("failed to unmarshal post, error: %v", err)
	}

	return account.Account.ID, post.Post.ID, nil
}
