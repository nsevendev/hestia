package controllers

import (
	depinject "hestia/app/depInject"
	"hestia/internal/logger"
	"hestia/internal/models"
	"hestia/internal/services"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

var newsService services.NewsService

func InitNewsController(c *depinject.Container) {
	newsService = c.NewsService
}

// GetAllNews récupère toutes les actualités
func GetAllNews(c *gin.Context) {
	statusCodeStr := c.Query("statusCode")
	statusCode, _ := strconv.Atoi(statusCodeStr)

	listNews, err := newsService.GetAll(c)
	if err != nil {
		c.HTML(statusCode, "private/news", gin.H{
			"Title":      "news",
			"Content":    "news",
			"ListNews":   nil,
			"Error":      "Erreur lors de la récupération des actualités",
			"Success":    nil,
		})
		return
	}

	success := c.Query("success")
	errorMsg := c.Query("error")

	c.HTML(http.StatusOK, "private/news", gin.H{
		"Title":      "news",
		"Content":    "news",
		"ListNews":   listNews,
		"Error":      errorMsg,
		"Success":    success,
	})
}

// affiche la liste + affiche info de la news selectionner dans le formulaire
func GetOneNews(c *gin.Context) {
	uuid := c.Params.ByName("uuid")

	listNews, errListNews := newsService.GetAll(c)
	if errListNews != nil {
		logger.Errorf("Erreur recuperation news / service-repository: %v", errListNews)
		c.Redirect(
			http.StatusSeeOther, 
			"/dashboard/news?statusCode=" + strconv.Itoa(http.StatusBadRequest) + "&error=" + url.QueryEscape("Erreur lors de la récupération des actualités"),
		)
		return
	}

	oneNewsMap := make(map[string]*models.News)
	for i := range listNews {
		news := &listNews[i]
		oneNewsMap[news.UUID.String()] = news
	}
	news, found := oneNewsMap[uuid]
	if !found {
		logger.Errorf("Erreur recuperation news / service-repository: erreur à la récuperation de l'actualité")
		c.Redirect(
			http.StatusSeeOther, 
			"/dashboard/news?statusCode=" + strconv.Itoa(http.StatusBadRequest) + "&error=" + url.QueryEscape("Erreur lors de la récupération de l'actualité"),
		)
		return
	}

	c.HTML(http.StatusOK, "private/news", gin.H{
		"Title":      "news",
		"Content":    "news",
		"ListNews":   listNews,
		"News":      news,
		"Error":      nil,
		"Success":    nil,
	})
}

// CreateNews crée une nouvelle actualité
func CreateNews(c *gin.Context) {
	titleRequest := c.PostForm("title")
	contentRequest := c.PostForm("content")
	linkTypeRequest := c.PostForm("linkType")
	logger.Infof("linkTypeRequest: %s", linkTypeRequest)
	title := validateDataStringEmpty(c, &titleRequest, "titre")
	content := validateDataStringEmpty(c, &contentRequest, "contenu")
	linkType := validateDataStringEmpty(c, &linkTypeRequest, "type de lien")

	var (
		linkFile *multipart.FileHeader
		linkURL  *string
	)

	if linkType == "file" {
		f, errFileLink := c.FormFile("link")
		if errFileLink != nil {
			logger.Errorf("Erreur upload file pour la creation news / lien requise: %v", errFileLink)
			c.Redirect(
				http.StatusSeeOther, 
				"/dashboard/news?statusCode=" + strconv.Itoa(http.StatusBadRequest) + "&error=" + url.QueryEscape("Erreur lors du téléchargement, lien requise"),
			)
			return
		}
		linkFile = f
	}

	if linkType == "url" {
		urlRequest := c.PostForm("link")
		url := validateDataStringEmpty(c, &urlRequest, "lien")
		linkURL = &url
	}
	
	fileImage, errFileImage := c.FormFile("image")
	if errFileImage != nil {
		logger.Errorf("Erreur upload file pour la creation news / image requise: %v", errFileImage)
		c.Redirect(
			http.StatusSeeOther, 
			"/dashboard/news?statusCode=" + strconv.Itoa(http.StatusBadRequest) + "&error=" + url.QueryEscape("Erreur lors du téléchargement, image requise"),
		)
		return
	}

	errCreate := newsService.Create(c.Request.Context(), &title, &content, fileImage, linkFile, linkURL, &linkType)
	if errCreate != nil {
		logger.Errorf("Erreur creation news / service-repository: %v", errCreate)
		c.Redirect(
			http.StatusSeeOther, 
			"/dashboard/news?statusCode=" + strconv.Itoa(http.StatusBadRequest) + "&error=" + url.QueryEscape("Erreur lors de la création de l'actualité, " + errCreate.Error()),
		)
		return
	}

	logger.Success("News à bien été créé")
	c.Redirect(
		http.StatusSeeOther, 
		"/dashboard/news?success=" + url.QueryEscape("Actualité créée avec succès"),
	)
}

// UpdateNews met à jour une actualité  existante
func UpdateNews(c *gin.Context) {
	// recuperer les donner de la request
	titleRequest := c.PostForm("title")
	contentRequest := c.PostForm("content")
	linkTypeRequest := c.PostForm("linkType")
	uuidNewsRequest := c.Params.ByName("uuid")
	title := validateDataStringEmpty(c, &titleRequest, "titre")
	content := validateDataStringEmpty(c, &contentRequest, "contenu")
	linkType := validateDataStringEmpty(c, &linkTypeRequest, "type de lien")
	uuidNews := validateDataStringEmpty(c, &uuidNewsRequest, "id de l'actualité")

	var (
		linkFile *multipart.FileHeader
		linkURL  *string
	)

	if linkType == "file" {
		f, errFileLink := c.FormFile("link")
		if errFileLink != nil {
			logger.Errorf("Erreur upload file pour la modification news / lien requis: %v", errFileLink)
			c.Redirect(
				http.StatusSeeOther, 
				"/dashboard/news?statusCode=" + strconv.Itoa(http.StatusBadRequest) + "&error=" + url.QueryEscape("Erreur lors du téléchargement, lien requis"),
			)
			return
		}
		linkFile = f
	}

	if linkType == "url" {
		urlRequest := c.PostForm("link")
		url := validateDataStringEmpty(c, &urlRequest, "lien")
		linkURL = &url
	}
	
	fileImage, errFileImage := c.FormFile("image")
	if errFileImage != nil {
		logger.Infof("Warning upload file pour la modifcation news / image n'existe pas et est non requise: %v", errFileImage)
	}

	errCreate := newsService.Update(c.Request.Context(), &title, &content, fileImage, linkFile, linkURL, &linkType, &uuidNews)
	if errCreate != nil {
		logger.Errorf("Erreur modification news / service-repository: %v", errCreate)
		c.Redirect(
			http.StatusSeeOther, 
			"/dashboard/news?statusCode=" + strconv.Itoa(http.StatusBadRequest) + "&error=" + url.QueryEscape("Erreur lors de la modification de l'actualité, " + errCreate.Error()),
		)
		return
	}

	logger.Success("News à bien été modifier")
	c.Redirect(
		http.StatusSeeOther, 
		"/dashboard/news?success=" + url.QueryEscape("Actualité modifier avec succès"),
	)
}

// DeleteNews supprime une actualité
func DeleteNews(c *gin.Context) {
	uuidRequest := c.Params.ByName("uuid")
	uuid := validateDataStringEmpty(c, &uuidRequest, "uuid")
	
	err := newsService.Delete(c.Request.Context(), &uuid)
	if err != nil {
		logger.Errorf("Erreur suppression news / service-repository: %v", err)
		c.Redirect(
			http.StatusSeeOther, 
			"/dashboard/news?statusCode=" + strconv.Itoa(http.StatusBadRequest) + "&error=" + url.QueryEscape("Erreur lors de la suppression de l'actualité"),
		)
		return
	}

	c.Redirect(
		http.StatusSeeOther, 
		"/dashboard/news?success=" + url.QueryEscape("Actualité supprimée avec succès"),
	)
}

func validateDataStringEmpty(c *gin.Context , value *string, name string) string {
	if *value == "" {
		c.Redirect(
			http.StatusSeeOther, 
			"/dashboard/news?statusCode=" + strconv.Itoa(http.StatusBadRequest) + "&error=" + url.QueryEscape("Erreur de validation, " + name + " requis"),
		)
	}

	return  *value
}
