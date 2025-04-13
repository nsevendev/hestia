package gallerycontroller

import (
	"hestia/internal/logger"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (gc *galleryController) AddImage(c *gin.Context) {
	titleRequest := c.PostForm("title")
	title := validateDataStringEmpty(c, &titleRequest, "titre")

	image, errImage := c.FormFile("image")
	if errImage != nil {
		logger.Errorf("[newscontroller::Create] Erreur telechargement fichier Link : %v", errImage)
		c.Redirect(
			http.StatusSeeOther,
			"/dashboard/gallery?statusCode="+strconv.Itoa(http.StatusBadRequest)+"&error="+url.QueryEscape("Erreur lors du téléchargement, image requise"),
		)
		return
	}

	if err := gc.galleryService.AddImage(c.Request.Context(), title, image); err != nil {
		logger.Errorf("[newscontroller::Create] Erreur creation Service: %v", err)
		c.Redirect(
			http.StatusSeeOther,
			"/dashboard/gallery?statusCode="+strconv.Itoa(http.StatusBadRequest)+"&error="+url.QueryEscape("Erreur lors de la création de l'actualité, "+err.Error()),
		)
		return
	}

	c.Redirect(
		http.StatusSeeOther,
		"/dashboard/gallery?statusCode="+strconv.Itoa(http.StatusOK)+"&success="+url.QueryEscape("L'image a correctement été ajoutée dans la galerie."),
	)
}
