package newscontroller

import (
	"hestia/internal/logger"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (nc newsController) Create(c *gin.Context) {
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
			logger.Errorf("[newscontroller::Create] Erreur telechargement fichier Link : %v", errFileLink)
			c.Redirect(
				http.StatusSeeOther,
				"/dashboard/news?statusCode="+strconv.Itoa(http.StatusBadRequest)+"&error="+url.QueryEscape("Erreur lors du téléchargement, lien requise"),
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
		logger.Errorf("[newscontroller::Create] Erreur telechargement fichier Image: %v", errFileImage)
		c.Redirect(
			http.StatusSeeOther,
			"/dashboard/news?statusCode="+strconv.Itoa(http.StatusBadRequest)+"&error="+url.QueryEscape("Erreur lors du téléchargement, image requise"),
		)
		return
	}

	if err := nc.newsService.Create(c.Request.Context(), &title, &content, fileImage, linkFile, linkURL, &linkType); err != nil {
		logger.Errorf("[newscontroller::Create] Erreur creation Service: %v", err)
		c.Redirect(
			http.StatusSeeOther,
			"/dashboard/news?statusCode="+strconv.Itoa(http.StatusBadRequest)+"&error="+url.QueryEscape("Erreur lors de la création de l'actualité, "+err.Error()),
		)
		return
	}

	c.Redirect(
		http.StatusSeeOther,
		"/dashboard/news?success="+url.QueryEscape("Actualité créée avec succès"),
	)
}
