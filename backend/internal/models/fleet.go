package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Fleet struct {
	ID           uuid.UUID `gorm:"type:char(36);primary_key" json:"id"`
	Name         string    `gorm:"type:varchar(100);not null" json:"name"`
	PlateNumber  string    `gorm:"type:varchar(20);uniqueIndex;not null" json:"plate_number"`
	Type         string    `gorm:"type:varchar(50);not null" json:"type"` // e.g., Hiace, Elf
	TotalSeats   int       `gorm:"not null" json:"total_seats"`
	ActiveStatus bool      `gorm:"default:true" json:"active_status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (f *Fleet) BeforeCreate(tx *gorm.DB) (err error) {
	if f.ID == uuid.Nil {
		f.ID = uuid.New()
	}
	return
}
