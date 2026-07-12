package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Route struct {
	ID              uuid.UUID `gorm:"type:char(36);primary_key"`
	OriginCity      string    `gorm:"type:varchar(100);index;not null"`
	OriginPool      string    `gorm:"type:varchar(150);not null"`
	DestinationCity string    `gorm:"type:varchar(100);index;not null"`
	DestinationPool string    `gorm:"type:varchar(150);not null"`
	CreatedAt       time.Time
}

func (r *Route) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return
}
