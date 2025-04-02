package news

import (
	"context"
	"hestia/internal/mediauri"
	"hestia/internal/models"
	"hestia/internal/upload"
	"mime/multipart"
	"path"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

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
func (s *newsService) Update(
	ctx context.Context, 
	title *string, 
	content *string, 
	image *multipart.FileHeader, 
	link *multipart.FileHeader, 
	url *string, 
	linkType *string, 
	uuidNews *string,
) error {
	news, errFindBy := s.GetById(ctx, uuidNews)
	if errFindBy != nil {
		return errFindBy
	}

	// ╔═══════════════════════════════════════════════════════════╗
	// ║                  UPDATE IMAGE FOR NEWS                    ║
	// ╚═══════════════════════════════════════════════════════════╝

	if image != nil {
		pathMediaImage := path.Join(s.pathPrefix, news.MediaImage.Path)
		if err := upload.DeleteFile(pathMediaImage); err != nil {
			return err
		}

		fileNameImage := uuid.New().String() + filepath.Ext(image.Filename)
		pathFinal := filepath.Join(s.pathBaseNews, s.folderForFile, "images", fileNameImage)
		if err := upload.SaveOrOverwrite(image,pathFinal); err != nil {
			return err
		}

		pathImage := upload.RemovePrefixPath(pathFinal, s.pathPrefix)
		if err := s.db.Model(news.MediaImage).Update("path", pathImage).Error; err != nil {
			return err
		}
	}

	// ╔═══════════════════════════════════════════════════════════╗
	// ║                  UPDATE LINK FOR NEWS                     ║
	// ╚═══════════════════════════════════════════════════════════╝

	if link != nil || url != nil {
		if news.MediaLink == nil {
			var mediaLink *models.MediaURI
			if *linkType == "file" {
				ext := filepath.Ext(link.Filename)
				typ, err := mediauri.DefineTypeFileMedia(ext)
				if err != nil {
					return err
				}

				fileNameLink := uuid.New().String() + filepath.Ext(link.Filename)
				pathFinal := filepath.Join(s.pathBaseNews, s.folderForFile, typ, fileNameLink)
				if err := upload.SaveOrOverwrite(link, pathFinal); err != nil {
					return err
				}
				pathLink := upload.RemovePrefixPath(pathFinal, s.pathPrefix)

				mediaLink = &models.MediaURI{
					UUID: uuid.New(),
					Path: pathLink,
					MediaType: typ,
					GalleryLinks: nil,
				}

				news.MediaLink = mediaLink
			}

			if *linkType == "url" {
				mediaLink = &models.MediaURI{
					UUID: uuid.New(),
					Path: *url,
					MediaType: "link",
					GalleryLinks: nil,
				}
				news.MediaLink = mediaLink
			}
		} else {
			if *linkType == "file" {
				ext := filepath.Ext(link.Filename)
				typ, err := mediauri.DefineTypeFileMedia(ext)
				if err != nil {
					return err
				}

				if news.MediaLink.MediaType == "audio" || news.MediaLink.MediaType == "pdf" || news.MediaLink.MediaType == "images" {
					pathMediaLink := path.Join(s.pathPrefix, news.MediaLink.Path)
					if err := upload.DeleteFile(pathMediaLink); err != nil {
						return err
					}
				}

				fileNameLink := uuid.New().String() + filepath.Ext(link.Filename)
				pathFinal := filepath.Join(s.pathBaseNews, s.folderForFile, typ, fileNameLink)
				if err := upload.SaveOrOverwrite(link, pathFinal); err != nil {
					return err
				}
				pathLink := upload.RemovePrefixPath(pathFinal, s.pathPrefix)

				if err := s.db.Model(news.MediaLink).Updates(map[string]any{
					"path":       pathLink,
					"media_type": typ,
				}).Error; err != nil {
					return err
				}
			}

			if *linkType == "url" {
				if news.MediaLink.MediaType == "audio" || news.MediaLink.MediaType == "pdf" || news.MediaLink.MediaType == "images" {
					pathMediaLink := path.Join(s.pathPrefix, news.MediaLink.Path)
					if err := upload.DeleteFile(pathMediaLink); err != nil {
						return err
					}
				}

				if err := s.db.Model(news.MediaLink).Updates(map[string]any{
					"path":       *url,
					"media_type": "link",
				}).Error; err != nil {
					return err
				}
			}
		}
	}

	// ╔═══════════════════════════════════════════════════════════╗
	// ║                        UPDATE NEWS                        ║
	// ╚═══════════════════════════════════════════════════════════╝

	news.UpdatedAt = time.Now()
	news.Title = *title
	news.Content = *content
	
	
	if err := s.db.WithContext(ctx).Save(news).Error; err != nil {
		return err
	}
	
	return nil
}