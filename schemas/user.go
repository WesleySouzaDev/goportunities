package schemas

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nome  string
	Email string
	Senha string
}

type UserResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Nome      string    `json:"nome"`
	Email     string    `json:"email"`
	Senha     string    `json:"senha"`
}
