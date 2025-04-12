package gallery

import (
	"context"
	"hestia/internal/models"
	"mime/multipart"

	"gorm.io/gorm"
)

// ╔═══════════════════════════════════════════════════════════╗
// ║                            PRIVATE                        ║
// ╚═══════════════════════════════════════════════════════════╝

type galleryService struct {
	db *gorm.DB
	pathBaseNews string
	pathPrefix  string
	folderForFile string
}	

// ╔═══════════════════════════════════════════════════════════╗
// ║                            PUBLIC                         ║
// ╚═══════════════════════════════════════════════════════════╝

type GalleryService interface {
	GetFirst() (*models.Gallery, error)
	AddImage(ctx context.Context, title string, image *multipart.FileHeader) error
	DeleteImageById(ctx context.Context, uuid string) error
}

func NewGalleryService(db *gorm.DB) GalleryService {
	return &galleryService{
		db,
		"app/views/assets/upload/",
		"app/views",
		"gallery",
	}
}