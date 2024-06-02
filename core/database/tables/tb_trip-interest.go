package tables

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (TripInterest) TableName() string {
	return tripInterestTableName
}

type TripInterest struct {
	TripId     uuid.UUID `json:"tripId"`
	InterestId uuid.UUID `json:"interestId"`

	CreatedDate     *time.Time `json:"createdDate"`
	IsDeleted       bool
	DeletedDate     *time.Time
	LastUpdatedDate *time.Time
}

func (ti *TripInterest) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	ti.CreatedDate = &now
	ti.IsDeleted = false
	return
}

func (ti *TripInterest) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	ti.LastUpdatedDate = &now
	return
}
