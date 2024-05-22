package tables

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (User) TableName() string {
	return "user_tb"
}

type User struct {
	Id           uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
	Username     string     `json:"username" gorm:"unique"`
	Password     string     `json:"password"`
	Name         string     `json:"name"`
	Surname      string     `json:"surname"`
	Email        string     `json:"email"`
	Gender       string     `json:"gender"`
	CreatedDate  *time.Time `json:"createdDate"`
	Biography    string     `json:"biography"`
	ProfileImage string     `json:"profileImage"`

	IsDeleted       bool
	DeletedDate     *time.Time
	LastUpdatedDate *time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	u.CreatedDate = &now
	u.IsDeleted = false
	u.Id = uuid.New()
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	u.LastUpdatedDate = &now
	return
}
