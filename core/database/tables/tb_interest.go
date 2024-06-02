package tables

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (Interest) TableName() string {
	return interestTableName
}

type Interest struct {
	Id          uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
	Name        string     `json:"name"`
	CreatedDate *time.Time `json:"createdDate"`

	IsDeleted       bool
	DeletedDate     *time.Time
	LastUpdatedDate *time.Time
}

func (u *Interest) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	u.CreatedDate = &now
	u.IsDeleted = false
	u.Id = uuid.New()
	return
}

func (u *Interest) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	u.LastUpdatedDate = &now
	return
}
