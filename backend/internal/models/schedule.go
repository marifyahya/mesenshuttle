package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Schedule struct {
	ID            uuid.UUID `gorm:"type:char(36);primary_key"`
	RouteID       uuid.UUID `gorm:"type:char(36);index;not null"`
	FleetID       uuid.UUID `gorm:"type:char(36);index;not null"`
	DepartureTime time.Time `gorm:"index;not null"`
	Price         int       `gorm:"not null"`
	CreatedAt     time.Time

	Route Route `gorm:"foreignKey:RouteID"`
	Fleet Fleet `gorm:"foreignKey:FleetID"`
}

func (s *Schedule) BeforeCreate(tx *gorm.DB) (err error) {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return
}
