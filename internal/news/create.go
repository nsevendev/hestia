package news

import (
	"context"
	"hestia/internal/mediauri"
	"hestia/internal/models"
	"hestia/internal/upload"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

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
func (s *newsService) Create(
	ctx context.Context, 
	title *string, 
	content *string, 
	image *multipart.FileHeader, 
	link *multipart.FileHeader, 
	url *string, 
	linkType *string,
) error {
	
	// ╔═══════════════════════════════════════════════════════════╗
	// ║                  CREATE LINK FOR NEWS                     ║
	// ╚═══════════════════════════════════════════════════════════╝
	
	var mediaLink *models.MediaURI

	if *linkType == "file" {
		ext := filepath.Ext(link.Filename)
		typ, errType := mediauri.DefineTypeFileMedia(ext)
		if errType != nil {
			return errType
		}
		
		customPathLink := filepath.Join(s.folderForFile, typ)
		fileNameLink := uuid.New().String() + ext
		pathLinkFinal, errLink := upload.PrepareFilePath(s.pathBaseNews, customPathLink, fileNameLink)
		if errLink != nil {
			return errLink
		}

		errUpload := upload.SaveOrOverwrite(link, pathLinkFinal)
		if errUpload != nil {
			return errUpload
		}

		mediaLink = &models.MediaURI{
			UUID: uuid.New(),
			Path: upload.RemovePrefixPath(pathLinkFinal, s.pathPrefix),
			MediaType: typ,
			GalleryLinks: nil,
		}
	}

	if *linkType == "url" {
		mediaLink = &models.MediaURI{
			UUID: uuid.New(),
			Path: *url,
			MediaType: "link",
			GalleryLinks: nil,
		}
	}

	// ╔═══════════════════════════════════════════════════════════╗
	// ║                  CREATE IMAGE FOR NEWS                    ║
	// ╚═══════════════════════════════════════════════════════════╝
	
	typ := "images"
	ext := filepath.Ext(image.Filename)
	customPathImage := filepath.Join(s.folderForFile, typ)
	fileNameImage := uuid.New().String() + ext
	pathImageFinal, errImage := upload.PrepareFilePath(s.pathBaseNews, customPathImage, fileNameImage)
	if errImage != nil {
		return errImage
	}

	errUpload := upload.SaveOrOverwrite(image, pathImageFinal)
	if errUpload != nil {
		return errUpload
	}

	mediaImage := &models.MediaURI{
		UUID: uuid.New(),
		Path: upload.RemovePrefixPath(pathImageFinal, s.pathPrefix),
		MediaType: typ,
		GalleryLinks: nil,
	}

	// ╔═══════════════════════════════════════════════════════════╗
	// ║                        CREATE NEWS                        ║
	// ╚═══════════════════════════════════════════════════════════╝

	news := &models.News{
		UUID: uuid.New(),
		Title: *title,
		Content: *content,
		PublishedAt: time.Now(),
		MediaImage: mediaImage,
		MediaLink: mediaLink,
	}

	err := s.db.WithContext(ctx).Create(news).Error
	if err != nil {
		return err
	}

	return nil
}