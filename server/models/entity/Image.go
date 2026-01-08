package entity

import (
	"server/models/dto"
	"time"
)

type Image struct {
	ID int `json:"id" gorm:"primaryKey"`
	dto.ImageDto
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
