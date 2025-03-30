package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email          string    `gorm:"uniqueIndex;not null"`
	HashedPassword string    `gorm:"not null"`
	Username 		string    `gorm:"uniqueIndex;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}