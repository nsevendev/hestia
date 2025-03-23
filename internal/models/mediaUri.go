package models

import (
	"time"

	"github.com/google/uuid"
)

type MediaURI struct {
	UUID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Path       string    `gorm:"not null"`
	MediaType  string    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	// Relations
	UsedInNewsImage []News `gorm:"foreignKey:UUIDMediaImage"`
	UsedInNewsLink  []News `gorm:"foreignKey:UUIDMediaLink"`
	GalleryLinks    []GalleryMediaURILink `gorm:"foreignKey:UUIDMediaURI"`
}