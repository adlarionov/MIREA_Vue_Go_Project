package entity

import (
	"server/models/dto"
	"time"
)

type Organization struct {
	ID uint `json:"id" gorm:"primaryKey"`
	dto.RegisterRequestDto
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Events    []Event   `json:"-"`
}
