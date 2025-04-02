package newscontroller

import (
	"hestia/internal/logger"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (nc newsController) DeleteById(c *gin.Context) {
	uuidRequest := c.Params.ByName("uuid")
	uuid := validateDataStringEmpty(c, &uuidRequest, "uuid")
	
	if err := nc.newsService.Delete(c.Request.Context(), &uuid); err != nil {
		logger.Errorf("[newscontroller::DeleteById]Erreur suppression Service : %v", err)
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