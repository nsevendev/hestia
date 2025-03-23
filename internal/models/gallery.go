package models

import (
	"time"

	"github.com/google/uuid"
)

type Gallery struct {
	UUID    uuid.UUID `gorm:"type:uuid;primaryKey"`
	Title   string    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Medias  []GalleryMediaURILink `gorm:"foreignKey:UUIDGallery"`
}