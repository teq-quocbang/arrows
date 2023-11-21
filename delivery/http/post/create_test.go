package post

import (
	"bytes"
	"encoding/json"
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

func TestCreate(t *testing.T) {
	assertion := assert.New(t)
	db := database.InitDatabase()
	defer db.TruncateTables()

	repo := repository.New(db.GetClient)
	r := Route{
		UseCase: usecase.New(repo, nil),
	}

	accountID, err := SetUpForeignKeyData(db)
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
		req := &payload.CreatePostRequest{
			Content:     gofakeit.BookGenre(),
			PrivacyMode: int32(proto.Privacy_Public.Number()),
		}
		resp, ctx := setupCreate(req)
		ctx.Set(string(auth.UserPrincipleKey), userPrinciple)

		// Act
		err = r.Create(ctx)

		// Assert
		assertion.NoError(err)
		assertion.Equal(200, resp.Code)
	}

	// bad case
	{
		// Arrange
		req := &payload.CreatePostRequest{}
		resp, ctx := setupCreate(req)

		// Act
		r.Create(ctx)

		// Assert
		assertion.Equal(400, resp.Code)
	}
}

func SetUpForeignKeyData(db *database.Database) (uuid.UUID, error) {
	repo := repository.New(db.GetClient)
	rAccount := account.Route{
		UseCase: usecase.New(repo, nil),
	}

	resp, ctx := setUpTestSignUp(&payload.SignUpRequest{
		Username: gofakeit.Name(),
		Password: gofakeit.Password(true, true, true, false, false, 10),
		Email:    gofakeit.Email(),
	})
	err := rAccount.SignUp(ctx)
	if err != nil {
		return uuid.UUID{}, err
	}
	if resp.Code != 200 {
		return uuid.UUID{}, fmt.Errorf("failed to sign up, error: %v", err)
	}
	account, err := test.UnmarshalBody[*presenter.AccountResponseWrapper](resp.Body.Bytes())
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to unmarshal sign up body, error: %v", err)
	}
	return account.Account.ID, nil
}

func setUpTestSignUp(input *payload.SignUpRequest) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	b, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "/api/user/sign-up", bytes.NewReader(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	return rec, c
}

func setupCreate(input *payload.CreatePostRequest) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	b, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "/api/post", bytes.NewReader(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	return rec, c
}
