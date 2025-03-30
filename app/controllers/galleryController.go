package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllNews récupère toutes les actualités
func GetGallery(c *gin.Context) {

	// appelle toute les news

	c.HTML(http.StatusOK, "private/gallery", gin.H{
		"Title":   "gallery",
		"Content": "gallery",
	})
}
