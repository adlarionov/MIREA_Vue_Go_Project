package entity

import (
	"server/models/dto"
	"time"
)

type Event struct {
	EventId uint `json:"id" gorm:"primaryKey"`
	dto.EventDto
	OwnerId   uint      `json:"user_id" gorm:"foreignKey:Users"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
