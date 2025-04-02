package depinject

import (
	"hestia/internal/news"

	"gorm.io/gorm"
)

type Container struct {
	NewsService news.NewsService
}

// NOTE : venir injecter ici les repositories et services
func NewContainer(db *gorm.DB) *Container {
	newsService := news.NewNewsService(db)

	return &Container{
		NewsService: newsService,
	}
}