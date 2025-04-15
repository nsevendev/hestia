package news

import (
	"context"
	"hestia/internal/models"
)

func (s *newsService) GetAll(ctx context.Context) ([]models.News, error) {
	var news []models.News

	err := s.db.WithContext(ctx).Preload("MediaImage").Preload("MediaLink").Order("created_at DESC").Find(&news).Error

	return news, err
}

func (s *newsService) GetById(ctx context.Context, uuid *string) (*models.News, error) {
	var news models.News

	err := s.db.WithContext(ctx).Preload("MediaImage").Preload("MediaLink").First(&news, "uuid = ?", uuid).Error
	if err != nil {
		return nil, err
	}

	return &news, nil
}
