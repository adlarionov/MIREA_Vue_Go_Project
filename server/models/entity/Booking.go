package entity

import (
	"server/models/dto"
	"time"
)

type Booking struct {
	ID int `json:"id" gorm:"primaryKey"`
	dto.BookingDto
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	EventId   *uint     `json:"-"`
	Event     Event     `gorm:"foreignKey:EventId"`
	VenueId   *uint     `json:"-"`
	Venue     Venue     `gorm:"foreignKey:VenueId"`
}

type BookingRequestDto struct {
	dto.BookingDto
	EventId *uint
	VenueId *uint
}

type BookingResponseDto struct {
	ID int `json:"id" gorm:"primaryKey"`
}
