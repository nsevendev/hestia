package services

import (
	"context"
	"fmt"
	"hestia/internal/logger"
	"hestia/internal/models"
	"hestia/internal/repository"
	"hestia/internal/upload"
	"mime/multipart"
	"path"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type NewsService interface {
	GetAll(ctx context.Context) ([]models.News, error)
	GetById(ctx context.Context, uuid *string) (*models.News, error)
	Create(ctx context.Context, title *string, content *string, image *multipart.FileHeader, link *multipart.FileHeader, url *string, linkType *string) error
	Update(ctx context.Context, title *string, content *string, image *multipart.FileHeader, link *multipart.FileHeader, url *string, linkType *string, uuidNews *string) error
	Delete(ctx context.Context, uuidNews *string) error
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

func (s *newsService) GetById(ctx context.Context, uuid *string) (*models.News, error) {
	return s.repo.FindById(ctx, uuid)
}

func (s *newsService) Create(ctx context.Context, title *string, content *string, image *multipart.FileHeader, link *multipart.FileHeader, url *string, linkType *string) error {
	/************************************ LINK NEWS (peut etre nul si pas renseingé) ************************************/

	var typ string
	var mediaLink *models.MediaURI

	switch *linkType {
		case "file":
			ext := filepath.Ext(link.Filename)
			switch ext {
				case ".mp3", ".wav", ".ogg", ".flac", ".mp4", ".m4a":
					typ = "audio"
				case ".pdf":
					typ = "pdf"
				case ".jepg", ".jpeg", ".png", ".gif", ".jpg":
					typ = "images"
				default:
					return fmt.Errorf("type de fichier non supporté, utiliser mp3, mp4, wav, pdf, jpg jepg, png, gif")
			}
			
			customPathLink := filepath.Join("news", typ)
	 		fileNameLink := uuid.New().String() + ext
			pathLink, err := upload.FileInUploadFolderWithCustomPath(link, &customPathLink, &fileNameLink, nil)
			if err != nil {
				return err
			}

			mediaLink = &models.MediaURI{
				UUID: uuid.New(),
				Path: pathLink,
				MediaType: typ,
				GalleryLinks: nil,
			}
		case "url":
			mediaLink = &models.MediaURI{
				UUID: uuid.New(),
				Path: *url,
				MediaType: "link",
				GalleryLinks: nil,
			}
		default:
			typ = ""
	}

	/******************** IMAGE NEWS (ne peux pas etre null obligatoire erreur dans le controller) **********************/

	customPathImage := filepath.Join("news", "images")
	fileNameImage := uuid.New().String() + filepath.Ext(image.Filename)
	pathImage, err := upload.FileInUploadFolderWithCustomPath(image, &customPathImage, &fileNameImage, nil)
	if err != nil {
		return err
	}

	mediaImage := &models.MediaURI{
		UUID: uuid.New(),
		Path: pathImage,
		MediaType: "images",
		GalleryLinks: nil,
	}

	/********************************************** CREATE NEWS *********************************************/

	news := &models.News{
		UUID: uuid.New(),
		Title: *title,
		Content: *content,
		PublishedAt: time.Now(),
		MediaImage: mediaImage,
		MediaLink: mediaLink,
	}

	return s.repo.Create(ctx, news)
}

func (s *newsService) Update(ctx context.Context, title *string, content *string, image *multipart.FileHeader, link *multipart.FileHeader, url *string, linkType *string, uuidNews *string) error {
	news, errFind := s.GetById(ctx, uuidNews)
	if errFind != nil {
		return errFind
	}

	/******************** IMAGE NEWS (peut etre nul si pas renseigné) **********************/

	if image != nil {
		pathMediaImage := path.Join("app/views", news.MediaImage.Path)
		errDeletePath := upload.DeleteFileInUploadFolder(&pathMediaImage)
		if errDeletePath != nil {
			return errDeletePath
		}

		customPathImage := filepath.Join("news", "images")
		fileNameImage := uuid.New().String() + filepath.Ext(image.Filename)
		pathImage, err := upload.FileInUploadFolderWithCustomPath(image, &customPathImage, &fileNameImage, nil)
		if err != nil {
			return err
		}

		errUpdate := s.repo.Db().Model(news.MediaImage).Update("path", pathImage).Error
		if errUpdate != nil {
			return errUpdate
		}
	}

	/************************************ LINK NEWS (peut etre nul si pas renseingé) ************************************/

	if link != nil || url != nil {
		if news.MediaLink == nil {
			var typp string
			var mediaLink *models.MediaURI
			logger.Infof("linkType: %s", *linkType)
			switch *linkType {
				case "file":
					ext := filepath.Ext(link.Filename)
					switch ext {
						case ".mp3", ".wav", ".ogg", ".flac", ".mp4", ".m4a":
							typp = "audio"
						case ".pdf":
							typp = "pdf"
						case ".jepg", ".jpeg", ".png", ".gif", ".jpg":
							typp = "images"
						default:
							return fmt.Errorf("type de fichier non supporté, utiliser mp3, mp4, wav, pdf, jpg jepg, png, gif")
					}
					
					customPathLink := filepath.Join("news", typp)
					fileNameLink := uuid.New().String() + ext
					pathLink, err := upload.FileInUploadFolderWithCustomPath(link, &customPathLink, &fileNameLink, nil)
					if err != nil {
						return err
					}

					mediaLink = &models.MediaURI{
						UUID: uuid.New(),
						Path: pathLink,
						MediaType: typp,
						GalleryLinks: nil,
					}

					news.MediaLink = mediaLink
				case "url":
					mediaLink = &models.MediaURI{
						UUID: uuid.New(),
						Path: *url,
						MediaType: "link",
						GalleryLinks: nil,
					}
					news.MediaLink = mediaLink
				default:
					typp = ""
			}
		} else {
			var typ string

			switch *linkType {
				case "file":
					ext := filepath.Ext(link.Filename)
					logger.Infof("ext: %s", ext)
					switch ext {
						case ".mp3", ".wav", ".ogg", ".flac", ".mp4", ".m4a":
							typ = "audio"
						case ".pdf":
							typ = "pdf"
						case ".jepg", ".jpeg", ".png", ".gif", ".jpg":
							typ = "images"
						default:
							return fmt.Errorf("type de fichier non supporté, utiliser mp3, mp4, wav, pdf, jpg jepg, jpeg, png, gif")
					}

					if news.MediaLink.MediaType == "audio" || news.MediaLink.MediaType == "pdf" || news.MediaLink.MediaType == "images" {
						pathMediaLink := path.Join("app/views", news.MediaLink.Path)
						errDeletePath := upload.DeleteFileInUploadFolder(&pathMediaLink)
						if errDeletePath != nil {
							return errDeletePath
						}
					}
					
					customPathLink := filepath.Join("news", typ)
					fileNameLink := uuid.New().String() + ext
					pathLink, err := upload.FileInUploadFolderWithCustomPath(link, &customPathLink, &fileNameLink, nil)
					if err != nil {
						return err
					}

					errUpdate := s.repo.Db().Model(news.MediaLink).Updates(map[string]any{
						"path":       pathLink,
						"media_type": typ,
					}).Error
					if errUpdate != nil {
						return errUpdate
					}
				case "url":
					if news.MediaLink.MediaType == "audio" || news.MediaLink.MediaType == "pdf" || news.MediaLink.MediaType == "images" {
						pathMediaLink := path.Join("app/views", news.MediaLink.Path)
						errDeletePath := upload.DeleteFileInUploadFolder(&pathMediaLink)
						if errDeletePath != nil {
							return errDeletePath
						}
					}

					errUpdate := s.repo.Db().Model(news.MediaLink).Updates(map[string]any{
						"path":       *url,
						"media_type": "link",
					}).Error
					if errUpdate != nil {
						return errUpdate
					}
				default:
					typ = ""
			}
		}
		
	}

	news.UpdatedAt = time.Now()
	news.Title = *title
	news.Content = *content
	
	return s.repo.Update(ctx, news)
}

func (s *newsService) Delete(ctx context.Context, uuidNews *string) error {
	news, errFind := s.GetById(ctx, uuidNews)
	if errFind != nil {
		return errFind
	}

	pathMediaImage := path.Join("app/views", news.MediaImage.Path)

	var pathMediaLink string
	if news.MediaLink != nil && (news.MediaLink.MediaType == "audio" || news.MediaLink.MediaType == "pdf" || news.MediaLink.MediaType == "images") {
		pathMediaLink = path.Join("app/views", news.MediaLink.Path)
	}

	errDelete := s.repo.Delete(ctx, news)
	if errDelete != nil {
		return errDelete
	}

	errDeletePath := upload.DeleteFileInUploadFolder(&pathMediaImage)
	if errDeletePath != nil {
		return errDeletePath
	}

	if pathMediaLink != "" {
		errDeletePathLink := upload.DeleteFileInUploadFolder(&pathMediaLink)
		if errDeletePathLink != nil {
			return errDeletePathLink
		}	
	}

	return nil
}