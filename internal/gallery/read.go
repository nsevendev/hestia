package gallery

import "hestia/internal/models"

func (s *galleryService) GetFirst() (*models.Gallery, error) {
	var gallery models.Gallery
    if err := s.db.Preload("Medias").Preload("Medias.MediaURI").First(&gallery).Error; err != nil {
        return nil, err
    }
    return &gallery, nil
}