package closureperiodcontroller

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (cc *closurePeriodController) Create(c *gin.Context) {
	title := c.PostForm("title")
	startDate := c.PostForm("startDate")
	endDate := c.PostForm("endDate")

	if err := cc.closureService.Create(c.Request.Context(), title, startDate, endDate) ;err != nil {
		c.Redirect(
			http.StatusSeeOther,
			"/dashboard/closure-period?statusCode=" + strconv.Itoa(http.StatusBadRequest) + "&error=" + url.QueryEscape("Erreur lors de la création de la période de fermeture, " + err.Error()),
		)
		return
	}

	c.Redirect(
		http.StatusSeeOther,
		"/dashboard/closure-period?statusCode=" + strconv.Itoa(http.StatusOK) + "&success=" + url.QueryEscape("Période de fermeture créée avec succès"),
	)
}