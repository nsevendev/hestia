package gallerycontroller

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (gc *galleryController) DeleteImageById(c *gin.Context) {
	uuid := c.Param("uuid")

	if err := gc.galleryService.DeleteImageById(c.Request.Context(), uuid); err != nil {
		c.Redirect(
			http.StatusSeeOther,
			"/dashboard/gallery?statusCode=" + strconv.Itoa(http.StatusBadRequest) + "&error=" + url.QueryEscape("Erreur lors de la suppression de l'image, de la gallery" + err.Error()),
		)
		return
	}

	c.Redirect(
		http.StatusSeeOther,
		"/dashboard/gallery?statusCode=" + strconv.Itoa(http.StatusOK) + "&success=" + url.QueryEscape("Image de la gallery supprimée avec succès"),
	)
}