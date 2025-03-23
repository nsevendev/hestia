package database

import (
	"hestia/internal/logger"
	"hestia/internal/models"
)

func AutoMigrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Term{},
		&models.MediaURI{},
		&models.ArticleTerm{},
		&models.News{},
		&models.Gallery{},
		&models.GalleryMediaURILink{},
	)

	if err != nil {
		logger.Fatalf("Erreur lors de la migration de la base de données : %v", err)
	}

	logger.Success("Migration de la base de données terminée avec succès.")
}