package entity

import (
	"server/models/dto"
	"time"
)

type Venue struct {
	ID int `json:"id" gorm:"primaryKey"`
	dto.VenueDto
	ImageId   *uint     `json:"-"`
	Image     Image     `gorm:"foreignKey:ImageId"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Bookings  []Booking `json:"-"`
}
