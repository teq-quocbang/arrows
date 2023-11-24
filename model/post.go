package model

import (
	"database/sql/driver"
	"encoding/json"
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/teq-quocbang/arrows/proto"
	"gorm.io/gorm"
)

type Emoji string

// emoji: [slice user id]
type React map[Emoji][]uuid.UUID

type PostInfo struct {
	// omitempty is important for JSON_MERGE_PATCH
	CommentIDs []uuid.UUID `json:"comment_ids,omitempty"`
	Reacts     React       `json:"reacts,omitempty"`
}

type Post struct {
	ID          uuid.UUID     `json:"id"`
	Content     string        `json:"content"`
	Information PostInfo      `json:"information"`
	PrivacyMode proto.Privacy `json:"privacy_mode"`
	CreatedAt   time.Time     `json:"created_at"`
	CreatedBy   uuid.UUID     `json:"created_by"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

func (m *PostInfo) Scan(src any) error {
	return ScanJSON(src, m)
}

func (m PostInfo) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (Post) TableName() string {
	return "post"
}

func (p *Post) ReactedThePost(userID uuid.UUID) (Emoji, bool) {
	for emoji, userIDs := range p.Information.Reacts {
		if slices.Contains[[]uuid.UUID](userIDs, userID) {
			return emoji, true
		}
	}
	return "", false
}

func (p *Post) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New()
	return nil
}

func (r React) ClearNilReact() {
	for e, userIDs := range r {
		if len(userIDs) == 0 {
			delete(r, e)
		}
	}
}
