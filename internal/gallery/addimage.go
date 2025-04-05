package gallery

import (
	"context"
	"hestia/internal/models"
	"hestia/internal/upload"
	"mime/multipart"
	"path/filepath"

	"github.com/google/uuid"
)

func (s *galleryService) AddImage(ctx context.Context, title string, image *multipart.FileHeader) error {
	galleryFirst, errGalleryFirst := s.GetFirst()
	if errGalleryFirst != nil {
		return errGalleryFirst
	}

	ext := filepath.Ext(image.Filename)
	customPathImage := filepath.Join(s.folderForFile, "images")
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
		MediaType: title,
		GalleryLinks: nil,
	}

	galleryMedia := &models.GalleryMediaURILink{
		UUID: uuid.New(),
		UUIDMediaURI: mediaImage.UUID,
		UUIDGallery: galleryFirst.UUID,
	}

	tx := s.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Create(&mediaImage).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&galleryMedia).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}