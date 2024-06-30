package tables

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (Like) TableName() string {
	return likeTableName
}

type Like struct {
	Id        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	PostId    uuid.UUID `json:"postId"`
	CommentId uuid.UUID `json:"commentId"`
	Username  string    `json:"userName"`

	CreatedDate *time.Time `json:"createdDate"`
	IsDeleted   bool       `json:"isDeleted"`
	DeletedDate *time.Time `json:"deletedDate"`
}

func (l *Like) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	l.CreatedDate = &now
	l.IsDeleted = false
	l.Id = uuid.New()
	return nil
}
