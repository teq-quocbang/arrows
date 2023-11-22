package post

import (
	"context"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/teq-quocbang/arrows/delivery/http/auth"
	"github.com/teq-quocbang/arrows/model"
	"github.com/teq-quocbang/arrows/proto"
	"github.com/teq-quocbang/arrows/repository/post"
	"github.com/teq-quocbang/arrows/util/myerror"
	"github.com/teq-quocbang/arrows/util/token"
)

func (s *TestSuite) TestGetByID() {
	assertion := s.Assertions
	userPrinciple := &token.JWTClaimCustom{
		User: token.UserInfo{
			ID: uuid.New(),
		},
	}
	postID := uuid.New()
	ctx := context.WithValue(s.ctx, auth.UserPrincipleKey, userPrinciple)

	// good case
	{
		// Arrange
		mockPost := post.NewMockRepository(s.T())
		mockPost.EXPECT().GetByID(ctx, postID).ReturnArguments = mock.Arguments{
			model.Post{
				ID:          postID,
				Content:     "test content",
				PrivacyMode: proto.Privacy_Public,
				CreatedBy:   userPrinciple.User.ID,
			}, nil,
		}
		u := s.useCase(mockPost)

		// Act
		reply, err := u.GetByID(ctx, postID)

		// Assert
		assertion.NoError(err)
		assertion.Equal(reply.Post.ID, postID)
	}

	// bad case
	{ // access denied
		// Arrange
		mockPost := post.NewMockRepository(s.T())
		mockPost.EXPECT().GetByID(ctx, postID).ReturnArguments = mock.Arguments{
			model.Post{
				ID:          postID,
				Content:     "test content",
				PrivacyMode: proto.Privacy_OnlyMe,
				CreatedBy:   uuid.New(),
			}, nil,
		}
		u := s.useCase(mockPost)

		// Act
		_, err := u.GetByID(ctx, postID)

		// Assert
		assertion.Error(err)
		assertion.Equal(myerror.ErrPostForbidden("access denied"), err)
	}
}
