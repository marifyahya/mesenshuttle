package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Schedule struct {
	ID            uuid.UUID `gorm:"type:char(36);primary_key" json:"id"`
	RouteID       uuid.UUID `gorm:"type:char(36);index;not null" json:"route_id"`
	Route         Route     `gorm:"foreignKey:RouteID" json:"route"`
	FleetID       uuid.UUID `gorm:"type:char(36);index;not null" json:"fleet_id"`
	Fleet         Fleet     `gorm:"foreignKey:FleetID" json:"fleet"`
	DepartureTime time.Time `gorm:"not null" json:"departure_time"`
	ArrivalTime   time.Time `gorm:"not null" json:"arrival_time"`
	Price         float64   `gorm:"type:decimal(10,2);not null" json:"price"`
	Status        string    `gorm:"type:varchar(20);default:'Scheduled'" json:"status"` // Scheduled, In Progress, Completed, Cancelled
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (s *Schedule) BeforeCreate(tx *gorm.DB) (err error) {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return
}
