package models

import (
	"time"

	"github.com/google/uuid"
)

type ClosurePeriod struct {
	UUID      uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title     string    `gorm:"size:255;not null"`         
	StartDate time.Time `gorm:"not null"`                  
	EndDate   time.Time `gorm:"not null"`                  
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (cp *ClosurePeriod) IsDateWithin(date time.Time) bool {
	return !date.Before(cp.StartDate) && !date.After(cp.EndDate)
}