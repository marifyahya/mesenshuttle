package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Fleet struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Type      string    `gorm:"type:varchar(50);not null"` // Premium/Reguler
	Capacity  int       `gorm:"not null"`
	CreatedAt time.Time
}

func (f *Fleet) BeforeCreate(tx *gorm.DB) (err error) {
	if f.ID == uuid.Nil {
		f.ID = uuid.New()
	}
	return
}
