package tables

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (Session) TableName() string {
	return "session_tb"
}

type Session struct {
	Id        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	UserId    uuid.UUID
	StartDate *time.Time
	EndDate   *time.Time

	CreatedDate *time.Time
	IsDeleted   bool
	DeletedDate *time.Time
}

func (u *Session) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	u.CreatedDate = &now
	u.StartDate = &now
	u.IsDeleted = false
	u.Id = uuid.New()

	return
}
