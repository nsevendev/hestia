package depinject

import (
	"hestia/internal/repository"
	"hestia/internal/services"

	"gorm.io/gorm"
)

type Container struct {
	NewsService services.NewsService
}

// NOTE : venir injecter ici les repositories et services
func NewContainer(db *gorm.DB) *Container {
	newsRepo := repository.NewNewsRepository(db)
	newsService := services.NewNewsService(newsRepo)

	return &Container{
		NewsService: newsService,
	}
}