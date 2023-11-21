package model

import (
	"database/sql/driver"
	"encoding/json"
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/teq-quocbang/store/proto"
	"gorm.io/gorm"
)

type emoji string

// emoji: [slice user id]
type React map[emoji][]uuid.UUID

type PostInfo struct {
	CommentIDs []uuid.UUID `json:"comment_ids"`
	Reacts     React       `json:"reacts"`
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

func (p *Post) ReactedThePost(userID uuid.UUID) (React, bool) {
	react := React{}
	for emoji, userIDs := range p.Information.Reacts {
		if slices.Contains[[]uuid.UUID](userIDs, userID) {
			react[emoji] = append(react[emoji], userID)
			return react, true
		}
	}
	return react, false
}

func (p *Post) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New()
	return nil
}
