package gallery

import (
	"context"
	"hestia/internal/models"
	"hestia/internal/upload"
	"path"
)

func (s *galleryService) DeleteImageById(ctx context.Context, uuid string) error {
	// recuperer gallerymedia, si existe ok sinon err
	tx := s.db.WithContext(ctx).Begin()
	if tx.Error != nil {
        return tx.Error
    }

	var media models.MediaURI
    if err := tx.First(&media, "uuid = ?", uuid).Error; err != nil {
        tx.Rollback()
        return err
    }

	pathMediaImage := path.Join(s.pathPrefix, media.Path)
	if err := upload.DeleteFile(pathMediaImage); err != nil {
        tx.Rollback()
		return err
	}

	if err := tx.Where("uuid_media_uri = ?", uuid).Delete(&models.GalleryMediaURILink{}).Error; err != nil {
        tx.Rollback()
        return err
    }

	if err := tx.Delete(&media).Error; err != nil {
        tx.Rollback()
        return err
    }

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}