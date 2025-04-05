package depinject

import (
	"hestia/internal/gallery"
	"hestia/internal/news"

	"gorm.io/gorm"
)

type Container struct {
	NewsService news.NewsService
	GalleryService gallery.GalleryService
}

// NOTE : venir injecter ici les repositories et services
func NewContainer(db *gorm.DB) *Container {
	newsService := news.NewNewsService(db)
	galleryService := gallery.NewGalleryService(db)

	return &Container{
		NewsService: newsService,
		GalleryService: galleryService,
	}
}