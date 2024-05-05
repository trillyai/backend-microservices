package tables

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (Post) TableName() string {
	return "post_tb"
}

type Post struct {
	Id          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	UserId      uuid.UUID `json:"userId"`
	TripId      uuid.UUID `json:"tripId"`
	Description string    `json:"description"`

	CreatedDate     *time.Time `json:"createdDate"`
	LastUpdatedDate *time.Time `json:"lastUpdateDate"`
	IsDeleted       bool       `json:"isDeleted"`
	DeletedDate     *time.Time `json:"deletedDate"`
}

func (p *Post) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	p.CreatedDate = &now
	p.IsDeleted = false
	p.Id = uuid.New()
	return nil
}

func (p *Post) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	p.LastUpdatedDate = &now
	return
}