package newscontroller

import (
	"hestia/internal/logger"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (nc *newsController) UpdateById(c *gin.Context) {
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
			logger.Errorf("[newscontroller::UpdateById] Erreur telechargement fichier pour la modification Link: %v", errFileLink)
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
		logger.Infof("[newscontroller::UpdateById] Attention telechargement fichier pour la modifcation Image : %v", errFileImage)
	}

	if err := nc.newsService.Update(c.Request.Context(), &title, &content, fileImage, linkFile, linkURL, &linkType, &uuidNews); err != nil {
		logger.Errorf("[newscontroller::UpdateById] Erreur modification Service : %v", err)
		c.Redirect(
			http.StatusSeeOther, 
			"/dashboard/news?statusCode=" + strconv.Itoa(http.StatusBadRequest) + "&error=" + url.QueryEscape("Erreur lors de la modification de l'actualité, " + err.Error()),
		)
		return
	}

	c.Redirect(
		http.StatusSeeOther, 
		"/dashboard/news?success=" + url.QueryEscape("Actualité modifier avec succès"),
	)
}