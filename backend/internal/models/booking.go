package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Booking struct {
	ID          uuid.UUID `gorm:"type:char(36);primary_key"`
	ScheduleID  uuid.UUID `gorm:"type:char(36);index;not null"`
	BookingCode string    `gorm:"type:varchar(20);uniqueIndex;not null"`
	Email       string    `gorm:"type:varchar(100);index;not null"`
	Name        string    `gorm:"type:varchar(100);not null"`
	Phone       string    `gorm:"type:varchar(20);index;not null"`
	Status      int8      `gorm:"type:tinyint;index;not null;default:0"` // 0:PENDING, 1:PAID, 2:EXPIRED
	TotalPrice  int       `gorm:"not null"`
	PaymentURL  string    `gorm:"type:varchar(255)"`
	CreatedAt   time.Time `gorm:"index"`

	Schedule     Schedule      `gorm:"foreignKey:ScheduleID"`
	BookingSeats []BookingSeat `gorm:"foreignKey:BookingID"`
}

type BookingSeat struct {
	ID         uuid.UUID `gorm:"type:char(36);primary_key"`
	BookingID  uuid.UUID `gorm:"type:char(36);index;not null"`
	SeatNumber string    `gorm:"type:varchar(10);not null"`
}

func (b *Booking) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}
	return
}

func (bs *BookingSeat) BeforeCreate(tx *gorm.DB) (err error) {
	if bs.ID == uuid.Nil {
		bs.ID = uuid.New()
	}
	return
}
