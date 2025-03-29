package services

import (
	"context"
	"hestia/internal/models"
	"hestia/internal/repository"
)

type NewsService interface {
	GetAll(ctx context.Context) ([]models.News, error)
	GetById(ctx context.Context, uuid string) (*models.News, error)
}

type newsService struct {
	repo repository.NewsRepository
}

func NewNewsService(repo repository.NewsRepository) NewsService {
	return &newsService{repo}
}

func (s *newsService) GetAll(ctx context.Context) ([]models.News, error) {
	return s.repo.FindAll(ctx)
}

func (s *newsService) GetById(ctx context.Context, uuid string) (*models.News, error) {
	return s.repo.FindById(ctx, uuid)
}