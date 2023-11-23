package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommentInfo struct {
	ChildComments []Comment `json:"child_comments"`
	Reacts        React     `json:"reacts"`
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
