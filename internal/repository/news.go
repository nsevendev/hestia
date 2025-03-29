package repository

import (
	"context"
	"hestia/internal/models"

	"gorm.io/gorm"
)

type NewsRepository interface {
	FindAll(ctx context.Context) ([]models.News, error)
	FindById(ctx context.Context, uuid string) (*models.News, error)
}

type newsRepository struct {
	db *gorm.DB
}

func NewNewsRepository(db *gorm.DB) NewsRepository {
	return &newsRepository{db}
}

func (r *newsRepository) FindAll(ctx context.Context) ([]models.News, error) {
	var news []models.News
	
	err := r.db.WithContext(ctx).Preload("MediaImage").Preload("MediaLink").Find(&news).Error

	return news, err
}

func (r *newsRepository) FindById(ctx context.Context, uuid string) (*models.News, error) {
	var news models.News

	err := r.db.WithContext(ctx).Preload("MediaImage").Preload("MediaLink").First(&news, "uuid = ?", uuid).Error
	if err != nil {
		return nil, err
	}
	
	return &news, nil
}