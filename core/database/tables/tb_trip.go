package tables

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (Trip) TableName() string {
	return tripTableName
}

type Trip struct {
	Id          uuid.UUID       `json:"id" gorm:"type:uuid;primaryKey"`
	DataJson    json.RawMessage `json:"data" gorm:"type:jsonb"`
	CreatedDate *time.Time      `json:"createdDate"`

	IsDeleted       bool
	DeletedDate     *time.Time
	LastUpdatedDate *time.Time
}

func (t *Trip) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	t.CreatedDate = &now
	t.IsDeleted = false
	t.Id = uuid.New()
	return
}

func (t *Trip) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	t.LastUpdatedDate = &now
	return
}
