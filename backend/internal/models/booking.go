package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Booking struct {
	ID              uuid.UUID     `gorm:"type:char(36);primary_key" json:"id"`
	ScheduleID      uuid.UUID     `gorm:"type:char(36);index;not null" json:"schedule_id"`
	BookingCode     string        `gorm:"type:varchar(20);uniqueIndex;not null" json:"booking_code"`
	Email           string        `gorm:"type:varchar(100);index;not null" json:"email"`
	Name            string        `gorm:"type:varchar(100);not null" json:"name"`
	Phone           string        `gorm:"type:varchar(20);index;not null" json:"phone"`
	Status          int8          `gorm:"type:tinyint;index;not null;default:0" json:"status"` // 0:PENDING, 1:PAID, 2:EXPIRED
	TotalPrice      int           `gorm:"not null" json:"total_price"`
	PaymentURL      string        `gorm:"type:varchar(255)" json:"payment_url"`
	CreatedAt       time.Time     `gorm:"index" json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`

	Schedule     Schedule      `gorm:"foreignKey:ScheduleID" json:"schedule"`
	BookingSeats []BookingSeat `gorm:"foreignKey:BookingID" json:"booking_seats"`
}

type BookingSeat struct {
	ID         uuid.UUID `gorm:"type:char(36);primary_key" json:"id"`
	BookingID  uuid.UUID `gorm:"type:char(36);index;not null" json:"booking_id"`
	SeatNumber string    `gorm:"type:varchar(10);not null" json:"seat_number"`
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
