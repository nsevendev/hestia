package news

import (
	"context"
	"hestia/internal/models"
)

/*
Récupère toutes les news

params:
	- ctx : context not null
return:
	- news : slice de news
	- error
*/
func (s *newsService) GetAll(ctx context.Context) ([]models.News, error) {
	var news []models.News
	
	err := s.db.WithContext(ctx).Preload("MediaImage").Preload("MediaLink").Find(&news).Error

	return news, err
}

/*
Récupère une news par son UUID

params:
	- ctx : context not null
	- uuid : pointeur string not null
return:
	- news : pointeur vers une news
	- error
*/
func (s *newsService) GetById(ctx context.Context, uuid *string) (*models.News, error) {
	var news models.News

	err := s.db.WithContext(ctx).Preload("MediaImage").Preload("MediaLink").First(&news, "uuid = ?", uuid).Error
	if err != nil {
		return nil, err
	}
	
	return &news, nil
}