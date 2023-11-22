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

func (s *TestSuite) TestUpsertEmoji() {
	assertion := s.Assertions

	userPrinciple := &token.JWTClaimCustom{
		User: token.UserInfo{
			ID: uuid.New(),
		},
	}
	ctx := context.WithValue(s.ctx, auth.UserPrincipleKey, userPrinciple)
	testPostID := uuid.New()
	testEmoji := gofakeit.Emoji()

	// good case
	{
		// Arrange
		mockPost := post.NewMockRepository(s.T())
		mockPost.EXPECT().GetByID(ctx, testPostID).ReturnArguments = mock.Arguments{
			model.Post{
				ID:          testPostID,
				Content:     gofakeit.BookGenre(),
				PrivacyMode: proto.Privacy_Public,
				CreatedBy:   userPrinciple.User.ID,
			}, nil,
		}
		react := &model.React{
			model.Emoji(testEmoji): []uuid.UUID{userPrinciple.User.ID},
		}
		mockPost.EXPECT().UpsertEmoji(ctx, testPostID, react).ReturnArguments = mock.Arguments{nil}
		u := s.useCase(mockPost)

		// Act
		reply, err := u.UpsertEmoji(ctx, &payload.UpsertEmojiRequest{
			PostID: testPostID.String(),
			Emoji:  testEmoji,
		})

		// Assert
		assertion.NoError(err)
		assertion.NotNil(reply)
	}

	// bad case
	{
		// Arrange
		u := s.useCase(post.NewMockRepository(s.T()))

		// Act
		_, err := u.UpsertEmoji(ctx, &payload.UpsertEmojiRequest{})

		// Assert
		assertion.Error(err)
	}
}
