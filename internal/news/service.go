package news

import (
	"context"
	"hestia/internal/models"
	"mime/multipart"

	"gorm.io/gorm"
)

type NewsService interface {
	Create(
		ctx context.Context, 
		title *string, 
		content *string, 
		image *multipart.FileHeader, 
		link *multipart.FileHeader, 
		url *string, 
		linkType *string,
	) error
	GetAll(ctx context.Context) ([]models.News, error)
	GetById(ctx context.Context, uuid *string) (*models.News, error)
	Update(
		ctx context.Context, 
		title *string, 
		content *string, 
		image *multipart.FileHeader, 
		link *multipart.FileHeader, 
		url *string, 
		linkType *string, 
		uuidNews *string,
	) error
	Delete(ctx context.Context, uuidNews *string) error
}

type newsService struct {
	pathBaseNews string
	pathPrefix  string
	folderForFile string
	db *gorm.DB
}

func NewNewsService(db *gorm.DB) NewsService {
	return &newsService{
		"app/views/assets/upload/",
		"app/views",
		"news",
		db,
	}
}