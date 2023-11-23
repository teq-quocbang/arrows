package model

import (
	"database/sql/driver"
	"encoding/json"
	"slices"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommentInfo struct {
	ChildCommentIDs []uuid.UUID `json:"child_comment_ids,omitempty"`
	ParentID        uuid.UUID   `json:"parent_id"`
	Reacts          React       `json:"reacts,omitempty"`
}

func (m *CommentInfo) Scan(src any) error {
	return ScanJSON(src, m)
}

func (m CommentInfo) Value() (driver.Value, error) {
	return json.Marshal(m)
}

type Comment struct {
	ID          uuid.UUID   `json:"id"`
	Contents    string      `json:"contents"`
	IsParent    bool        `json:"is_parent"`
	Information CommentInfo `json:"information"`
	CreatedAt   time.Time   `json:"created_at"`
	CreatedBy   uuid.UUID   `json:"created_by"`
	UpdatedAt   time.Time   `json:"updated_at"`
	PostID      uuid.UUID   `json:"post_id"`
}

func (c *Comment) ReactedThePost(userID uuid.UUID) (Emoji, bool) {
	for emoji, userIDs := range c.Information.Reacts {
		if slices.Contains[[]uuid.UUID](userIDs, userID) {
			return emoji, true
		}
	}
	return "", false
}

func (m *Comment) Scan(src any) error {
	return ScanJSON(src, m)
}

func (m Comment) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (Comment) TableName() string {
	return "comment"
}

func (c *Comment) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.New()
	return nil
}
