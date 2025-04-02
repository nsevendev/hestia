package news

import (
	"context"
	"hestia/internal/models"
	"hestia/internal/upload"
	"path"
)

/*
Supprime une news et ses fichiers associés (image, pdf, fichier audio)

params:
	- ctx : context not null
	- uuidNews : pointeur string not null
return:
	- error
*/
func (s *newsService) Delete(ctx context.Context, uuidNews *string) error {
	news, errFindById := s.GetById(ctx, uuidNews)
	if errFindById != nil {
		return errFindById
	}

	pathMediaImage := path.Join(s.pathPrefix, news.MediaImage.Path)

	// ╔═══════════════════════════════════════════════════════════╗
	// ║                      DELETE ALL MODELS                    ║
	// ╚═══════════════════════════════════════════════════════════╝

	if err := s.db.WithContext(ctx).Delete(&news.MediaImage).Error; err != nil {
		return err
	}

	var pathMediaLink string
	if news.MediaLink != nil {
		pathMediaLink = path.Join(s.pathPrefix, news.MediaLink.Path)
		if err := s.db.WithContext(ctx).Delete(&news.MediaLink).Error; err != nil {
			return err
		}	
	}

	if err := s.db.WithContext(ctx).Delete(&models.News{}, "uuid = ?", &news.UUID).Error; err != nil {
		return err
	}

	// ╔═══════════════════════════════════════════════════════════╗
	// ║                       DELETE FILES                        ║
	// ╚═══════════════════════════════════════════════════════════╝

	if err := upload.DeleteFile(pathMediaImage); err != nil {
		return err
	}

	if pathMediaLink != "" {
		if err := upload.DeleteFile(pathMediaLink); err != nil {
			return err
		}	
	}

	return nil
}