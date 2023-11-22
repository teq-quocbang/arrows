package post

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/teq-quocbang/arrows/delivery/http/auth"
	"github.com/teq-quocbang/arrows/fixture/database"
	"github.com/teq-quocbang/arrows/payload"
	"github.com/teq-quocbang/arrows/proto"
	"github.com/teq-quocbang/arrows/repository"
	"github.com/teq-quocbang/arrows/usecase"
	"github.com/teq-quocbang/arrows/util/token"
)

func TestUpsertEmoji(t *testing.T) {
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
		req := &payload.UpsertEmojiRequest{
			PostID: postID.String(),
			Emoji:  gofakeit.Emoji(),
		}
		resp, ctx := setupUpsertEmoji(req)
		ctx.Set(string(auth.UserPrincipleKey), userPrinciple)

		// Act
		err := r.UpsertEmoji(ctx)

		//Assert
		assertion.NoError(err)
		assertion.Equal(200, resp.Code)
	}

	// bad case
	{
		// Arrange
		resp, ctx := setupUpsertEmoji(nil)
		ctx.Set(string(auth.UserPrincipleKey), userPrinciple)

		// Act
		err := r.UpsertEmoji(ctx)

		//Assert
		assertion.NoError(err)
		assertion.Equal(400, resp.Code)
	}
}

func setupUpsertEmoji(input *payload.UpsertEmojiRequest) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	b, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPatch, "/api/post", bytes.NewReader(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	return rec, c
}

func TestUpdate(t *testing.T) {
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
	testOnlyMeMode := proto.Privacy_OnlyMe

	// good case
	{
		// Arrange
		req := &payload.UpdatePostRequest{
			PostID:      postID.String(),
			PrivacyMode: int32(testOnlyMeMode),
		}
		resp, ctx := setupUpdate(req)
		ctx.Set(string(auth.UserPrincipleKey), userPrinciple)

		// Act
		err := r.Update(ctx)

		// Assert
		assertion.NoError(err)
		assertion.Equal(200, resp.Code)
	}

	// good case
	{
		// Arrange
		resp, ctx := setupUpdate(nil)
		ctx.Set(string(auth.UserPrincipleKey), userPrinciple)

		// Act
		r.Update(ctx)

		// Assert
		assertion.Equal(400, resp.Code)
	}
}

func setupUpdate(input *payload.UpdatePostRequest) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	b, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPut, "/api/post", bytes.NewReader(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	return rec, c
}
