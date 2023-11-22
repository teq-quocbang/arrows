package post

import (
	"context"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/teq-quocbang/arrows/delivery/http/auth"
	"github.com/teq-quocbang/arrows/model"
	"github.com/teq-quocbang/arrows/payload"
	"github.com/teq-quocbang/arrows/proto"
	"github.com/teq-quocbang/arrows/repository/post"
	"github.com/teq-quocbang/arrows/util/token"
)

func (s *TestSuite) TestCreate() {
	assertion := s.Assertions
	userPrinciple := &token.JWTClaimCustom{
		SessionID: uuid.New(),
		User: token.UserInfo{
			ID: uuid.New(),
		},
	}

	ctx := context.WithValue(s.ctx, auth.UserPrincipleKey, userPrinciple)
	testPrivacyMode := proto.Privacy_Public
	content := gofakeit.BookGenre()
	// good case
	{
		// Arrange
		mockPost := post.NewMockRepository(s.T())
		postModel := &model.Post{
			Content:     content,
			PrivacyMode: testPrivacyMode,
			CreatedBy:   userPrinciple.User.ID,
		}
		mockPost.EXPECT().Create(ctx, postModel).ReturnArguments = mock.Arguments{nil}
		req := &payload.CreatePostRequest{
			Content:     content,
			PrivacyMode: int32(proto.Privacy_Public),
		}
		u := s.useCase(mockPost)

		// Act
		reply, err := u.Create(ctx, req)

		// Assert
		assertion.NoError(err)
		assertion.NotNil(reply)
	}
	// bad case
	{
		// Arrange
		u := s.useCase(post.NewMockRepository(s.T()))

		// Act
		_, err := u.Create(ctx, nil)

		// Assert
		assertion.Error(err)
	}
}
