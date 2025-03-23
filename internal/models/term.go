package models

import (
	"time"

	"github.com/google/uuid"
)

type Term struct {
	UUID      uuid.UUID `gorm:"type:uuid;primaryKey"`
	Title     string    `gorm:"not null"`
	SubTitle  string
	Articles  []ArticleTerm `gorm:"foreignKey:UUIDTerm"`
	CreatedAt time.Time
	UpdatedAt time.Time
}