package entity

import (
	"server/models/dto"
	"time"
)

const (
	RoleUser  string = "user"
	RoleAdmin string = "admin"
	RoleOwner string = "owner"
)

type User struct {
	UserId uint `json:"id" gorm:"primaryKey"`
	dto.UserRegisterDto
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
