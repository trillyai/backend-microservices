package tables

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (Venue) TableName() string {
	return venueTableName
}

type Venue struct {
	Id          uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
	Name        string     `json:"name"`
	Latitude    string     `json:"latitude"`
	Longitude   string     `json:"longitude"`
	CreatedDate *time.Time `json:"createdDate"`

	IsDeleted       bool
	DeletedDate     *time.Time
	LastUpdatedDate *time.Time
}

func (v *Venue) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	v.CreatedDate = &now
	v.IsDeleted = false
	v.Id = uuid.New()
	return
}

func (v *Venue) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	v.LastUpdatedDate = &now
	return
}
