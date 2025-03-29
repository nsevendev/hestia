package controllers

import (
	"net/http"
	"time"

	"hestia/internal/database"
	"hestia/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetAllNews récupère toutes les actualités
func GetAllNews(c *gin.Context) {
	var news []models.News
	if err := database.DB.Find(&news).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, news)
}

// CreateNews crée une nouvelle actualité
func CreateNews(c *gin.Context) {
	var news models.News
	if err := c.ShouldBindJSON(&news); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Générer un nouvel UUID et définir les timestamps
	news.UUID = uuid.New()
	news.CreatedAt = time.Now()
	news.UpdatedAt = time.Now()

	if err := database.DB.Create(&news).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, news)
}

// UpdateNews met à jour une actualité  existante
func UpdateNews(c *gin.Context) {
	idParam := c.Param("id")
	newsUUID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UUID invalide"})
		return
	}

	// Récupérer l'actualité existante
	var news models.News
	if err := database.DB.First(&news, "uuid = ?", newsUUID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Actualité non trouvée"})
		return
	}

	// Récupérer les données envoyées dans le body
	var updateData models.News
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Mise à jour des champs souhaités
	news.Title = updateData.Title
	news.Content = updateData.Content
	news.PublishedAt = updateData.PublishedAt
	news.UUIDMediaImage = updateData.UUIDMediaImage
	news.UUIDMediaLink = updateData.UUIDMediaLink
	news.UpdatedAt = time.Now()

	if err := database.DB.Save(&news).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, news)
}

// DeleteNews supprime une actualité
func DeleteNews(c *gin.Context) {
	idParam := c.Param("id")
	newsUUID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UUID invalide"})
		return
	}

	if err := database.DB.Delete(&models.News{}, "uuid = ?", newsUUID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Actualité supprimée avec succès"})
}
