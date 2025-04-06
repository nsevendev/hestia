package closureperiodcontroller

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (cc *closurePeriodController) DeleteById(c *gin.Context) {
	uuidClosurePeriod := c.Param("uuid")

	if err := cc.closureService.Delete(c.Request.Context(), uuidClosurePeriod); err != nil {
		c.Redirect(
			http.StatusSeeOther,
			"/dashboard/closure-period?statusCode=" + strconv.Itoa(http.StatusBadRequest) + "&error=" + url.QueryEscape("Erreur lors de la suppression de la période de fermeture, " + err.Error()),
		)
		return
	}

	c.Redirect(
		http.StatusSeeOther,
		"/dashboard/closure-period?statusCode=" + strconv.Itoa(http.StatusOK) + "&success=" + url.QueryEscape("Période de fermeture supprimée avec succès"),
	)
}