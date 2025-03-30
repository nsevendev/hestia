package repository

import (
	"context"
	"hestia/internal/models"

	"gorm.io/gorm"
)

type NewsRepository interface {
	FindAll(ctx context.Context) ([]models.News, error)
	FindById(ctx context.Context, uuid *string) (*models.News, error)
	Create(ctx context.Context, news *models.News) error
	Update(ctx context.Context, news *models.News) error
	Delete(ctx context.Context, news *models.News) error
	Db() *gorm.DB
}

type newsRepository struct {
	db *gorm.DB
}

func NewNewsRepository(db *gorm.DB) NewsRepository {
	return &newsRepository{db}
}

func (r *newsRepository) Db() *gorm.DB {
	return r.db
}

func (r *newsRepository) FindAll(ctx context.Context) ([]models.News, error) {
	var news []models.News
	
	err := r.db.WithContext(ctx).Preload("MediaImage").Preload("MediaLink").Find(&news).Error

	return news, err
}

func (r *newsRepository) FindById(ctx context.Context, uuid *string) (*models.News, error) {
	var news models.News

	err := r.db.WithContext(ctx).Preload("MediaImage").Preload("MediaLink").First(&news, "uuid = ?", uuid).Error
	if err != nil {
		return nil, err
	}
	
	return &news, nil
}

func (r *newsRepository) Create(ctx context.Context, news *models.News) error {
	err := r.db.WithContext(ctx).Create(news).Error
	if err != nil {
		return err
	}
	
	return nil
}

func (r *newsRepository) Update(ctx context.Context, news *models.News) error {
	err := r.db.WithContext(ctx).Save(news).Error
	if err != nil {
		return err
	}
	
	return nil
}

func (r *newsRepository) Delete(ctx context.Context, news *models.News) error {
	errMediaImage := r.db.WithContext(ctx).Delete(&news.MediaImage).Error
	if errMediaImage != nil {
		return errMediaImage
	}
	
	if news.MediaLink != nil {
		errMediaLink := r.db.WithContext(ctx).Delete(&news.MediaLink).Error
		if errMediaLink != nil {
			return errMediaLink
		}	
	}

	err := r.db.WithContext(ctx).Delete(&models.News{}, "uuid = ?", &news.UUID).Error
	if err != nil {
		return err
	}
	
	return nil
}