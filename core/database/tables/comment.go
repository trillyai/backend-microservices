package tables

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (Comment) TableName() string {
	return "comment_tb"
}

type Comment struct {
	Id      uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	PostId  uuid.UUID `json:"postId"`
	UserId  uuid.UUID `json:"userId"`
	Comment string    `json:"comment"`

	CreatedDate     *time.Time `json:"createdDate"`
	LastUpdatedDate *time.Time `json:"lastUpdateDate"`
	IsDeleted       bool       `json:"isDeleted"`
	DeletedDate     *time.Time `json:"deletedDate"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	c.CreatedDate = &now
	c.IsDeleted = false
	c.Id = uuid.New()
	return nil
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	c.LastUpdatedDate = &now
	return
}
