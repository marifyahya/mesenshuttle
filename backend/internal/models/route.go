package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Route struct {
	ID              uuid.UUID `gorm:"type:char(36);primary_key" json:"id"`
	OriginCity      string    `gorm:"type:varchar(100);index;not null" json:"origin_city"`
	OriginPool      string    `gorm:"type:varchar(150);not null" json:"origin_pool"`
	DestinationCity string    `gorm:"type:varchar(100);index;not null" json:"destination_city"`
	DestinationPool string    `gorm:"type:varchar(150);not null" json:"destination_pool"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (r *Route) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return
}
