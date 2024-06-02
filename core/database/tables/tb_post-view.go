package tables

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (PostView) TableName() string {
	return postViewTableName
}

type PostView struct {
	Username    string     `json:"username"`
	PostId      uuid.UUID  `json:"postId"`
	CreatedDate *time.Time `json:"createdDate"`

	IsDeleted       bool
	DeletedDate     *time.Time
	LastUpdatedDate *time.Time
}

func (pw *PostView) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	pw.CreatedDate = &now
	pw.IsDeleted = false
	return
}

func (pw *PostView) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	pw.LastUpdatedDate = &now
	return
}
