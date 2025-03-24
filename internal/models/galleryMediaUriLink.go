package models

import (
	"time"

	"github.com/google/uuid"
)

type GalleryMediaURILink struct {
	UUID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UUIDGallery   uuid.UUID `gorm:"type:uuid;not null"`
	UUIDMediaURI  uuid.UUID `gorm:"type:uuid;not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Gallery   Gallery  `gorm:"foreignKey:UUIDGallery;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	MediaURI  MediaURI `gorm:"foreignKey:UUIDMediaURI;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}