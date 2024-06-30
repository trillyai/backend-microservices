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

func (i *Interest) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	i.CreatedDate = &now
	i.IsDeleted = false
	i.Id = uuid.New()
	return
}

func (i *Interest) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	i.LastUpdatedDate = &now
	return
}
