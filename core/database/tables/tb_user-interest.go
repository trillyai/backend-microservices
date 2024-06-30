package tables

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (UserInterest) TableName() string {
	return userIntrestTableName
}

type UserInterest struct {
	UserName    string     `json:"username"`
	InterestId  uuid.UUID  `json:"interestId"`
	CreatedDate *time.Time `json:"createdDate"`

	IsDeleted       bool
	DeletedDate     *time.Time
	LastUpdatedDate *time.Time
}

func (u *UserInterest) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	u.CreatedDate = &now
	u.IsDeleted = false
	return
}

func (u *UserInterest) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	u.LastUpdatedDate = &now
	return
}
