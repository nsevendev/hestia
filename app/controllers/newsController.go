package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllNews récupère toutes les actualités
func GetAllNews(c *gin.Context) {

	// appelle toute les news


	c.HTML(http.StatusOK, "private/news", gin.H{
		"Title": "news",
		"Content": "news",
	})
}

// affiche la liste + affiche info de la news selectionner dans le formulaire
func GetOneNews(c *gin.Context) {
	// appelle toute les news

	// recuperation de la news avec l'id


	// return la liste des news
	// return la news selectionner

	// return news.html
}

// CreateNews crée une nouvelle actualité
func CreateNews(c *gin.Context) {
	// fait le traitement
	// redirection
}

// UpdateNews met à jour une actualité  existante
func UpdateNews(c *gin.Context) {
	// fait le traitement
	// redirection
}

// DeleteNews supprime une actualité
func DeleteNews(c *gin.Context) {
	// fait le traitement
	// redirection
}
