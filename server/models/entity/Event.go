package entity

import (
	"server/models/dto"
	"time"
)

type Event struct {
	ID int `json:"id" gorm:"primaryKey"`
	dto.EventDto
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
	OrganizationId uint         `json:"-"`
	Organization   Organization `gorm:"foreignKey:OrganizationId"`
}

type EventResponseDto struct {
	ID int `json:"id" gorm:"primaryKey"`
	dto.EventDto
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	Organization Organization `gorm:"foreignKey:OrganizationId"`
}
