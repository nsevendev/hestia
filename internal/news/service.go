package news

import (
	"context"
	"hestia/internal/models"
	"mime/multipart"

	"gorm.io/gorm"
)

// ╔═══════════════════════════════════════════════════════════╗
// ║                            PRIVATE                        ║
// ╚═══════════════════════════════════════════════════════════╝


type newsService struct {
	pathBaseNews string
	pathPrefix  string
	folderForFile string
	db *gorm.DB
}

// ╔═══════════════════════════════════════════════════════════╗
// ║                            PUBLIC                         ║
// ╚═══════════════════════════════════════════════════════════╝

type NewsService interface {
	/*
	Creation de la news

	params:
		- ctx : context not null
		- title : pointeur string not null
		- content : pointeur string not null
		- image : pointeur file not null
		- link : pointeur file nullable
		- url : pointeur string nullable
		- linkType : pointeur string empty ("file", "url", "")
	return:
		- error
	*/
	Create(
		ctx context.Context, 
		title *string, 
		content *string, 
		image *multipart.FileHeader, 
		link *multipart.FileHeader, 
		url *string, 
		linkType *string,
	) error
	
	/*
	Récupère toutes les news

	params:
		- ctx : context not null
	return:
		- news : slice de news
		- error
	*/
	GetAll(ctx context.Context) ([]models.News, error)
	
	/*
	Récupère une news par son UUID

	params:
		- ctx : context not null
		- uuid : pointeur string not null
	return:
		- news : pointeur vers une news
		- error
	*/
	GetById(ctx context.Context, uuid *string) (*models.News, error)
	
	/*
	modifie une news et ses fichiers associés (image, pdf, fichier audio)

	params:
		- ctx : context not null
		- title : pointeur string not null
		- content : pointeur string not null
		- image : pointeur file nullable
		- link : pointeur file nullable
		- url : pointeur string nullable
		- linkType : pointeur string empty ("file", "url", "")
		- uuidNews : pointeur string not null
	return:
		- error
	*/
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
	
	/*
	Supprime une news et ses fichiers associés (image, pdf, fichier audio)

	params:
		- ctx : context not null
		- uuidNews : pointeur string not null
	return:
		- error
	*/
	Delete(ctx context.Context, uuidNews *string) error
}

func NewNewsService(db *gorm.DB) NewsService {
	return &newsService{
		"app/views/assets/upload/",
		"app/views",
		"news",
		db,
	}
}

