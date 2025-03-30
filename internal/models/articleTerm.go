package models

import (
	"time"

	"github.com/google/uuid"
)

type ArticleTerm struct {
	UUID      uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title     string    `gorm:"not null"`
	Content   string    `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UUIDTerm  uuid.UUID `gorm:"type:uuid;not null"`
	Term      Term      `gorm:"foreignKey:UUIDTerm;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}