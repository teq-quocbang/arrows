package model

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestReactedThePost(t *testing.T) {
	assertion := assert.New(t)

	testUser := uuid.New()
	testEmoji := gofakeit.Emoji()
	// good case
	{
		// Arrange
		react := React{}
		for i := 0; i < 10; i++ {
			react[Emoji(gofakeit.Emoji())] = fakeUserIDs(20)
		}
		react[Emoji(testEmoji)] = []uuid.UUID{testUser} // for testing purpose
		post := Post{
			ID:      uuid.New(),
			Content: gofakeit.BookGenre(),
			Information: PostInfo{
				CommentIDs: []uuid.UUID{uuid.New()},
				Reacts:     react,
			},
		}

		// Act
		reply, ok := post.ReactedThePost(testUser)

		// Assert
		assertion.True(ok)
		assertion.Equal(Emoji(testEmoji), reply)
	}
}

func fakeUserIDs(in int) []uuid.UUID {
	result := make([]uuid.UUID, in)
	for i := 0; i < in; i++ {
		result[i] = uuid.New()
	}
	return result
}
