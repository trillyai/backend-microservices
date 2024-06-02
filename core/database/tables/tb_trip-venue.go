package tables

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (TripVenue) TableName() string {
	return tripVenueTableName
}

type TripVenue struct {
	TripId      uuid.UUID  `json:"tripId"`
	VenueId     uuid.UUID  `json:"venueId"`
	CreatedDate *time.Time `json:"createdDate"`

	IsDeleted       bool
	DeletedDate     *time.Time
	LastUpdatedDate *time.Time
}

func (tv *TripVenue) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	tv.CreatedDate = &now
	tv.IsDeleted = false
	return
}

func (tv *TripVenue) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	tv.LastUpdatedDate = &now
	return
}
