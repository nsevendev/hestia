package models

import (
	"time"

	"github.com/google/uuid"
)

type News struct {
	UUID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title           string    `gorm:"not null"`
	Content         string    `gorm:"type:text"`
	PublishedAt     time.Time
	UUIDMediaImage  uuid.UUID `gorm:"type:uuid"`
	UUIDMediaLink   uuid.UUID `gorm:"type:uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time

	MediaImage      *MediaURI `gorm:"foreignKey:UUIDMediaImage;constraint:OnUpdate:SET NULL,OnDelete:SET NULL"`
	MediaLink       *MediaURI `gorm:"foreignKey:UUIDMediaLink;constraint:OnUpdate:SET NULL,OnDelete:SET NULL"`
}