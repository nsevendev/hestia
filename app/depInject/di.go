package depinject

import (
	"hestia/internal/auth"
	"hestia/internal/closedperiod"
	"hestia/internal/gallery"
	"hestia/internal/news"

	"gorm.io/gorm"
)

type Container struct {
	NewsService news.NewsService
	GalleryService gallery.GalleryService
	AuthService auth.AuthService
	ClosedPeriodService closedperiod.ClosedPeriodService
}

// NOTE : venir injecter ici les repositories et services
func NewContainer(db *gorm.DB) *Container {
	authService := auth.NewAuthService(db)
	newsService := news.NewNewsService(db)
	galleryService := gallery.NewGalleryService(db)
	closurePeriod := closedperiod.NewClosedPeriodService(db)

	return &Container{
		NewsService: newsService,
		GalleryService: galleryService,
		AuthService: authService,
		ClosedPeriodService: closurePeriod,
	}
}