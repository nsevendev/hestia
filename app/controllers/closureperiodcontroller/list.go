package closureperiodcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (cc *closurePeriodController) List(c *gin.Context) {
	listPeriod, err := cc.closureService.List(c)
	if err != nil {
		cc.res.ListPeriod = nil
		cc.res.Error = "Erreur lors de la récupération des périodes de fermeture"
		c.HTML(http.StatusInternalServerError, "private/closure", cc.res)
		return
	}

	cc.res.ListPeriod = listPeriod

	cc.res.Success = c.Query("success")
	cc.res.Error = c.Query("error")

	c.HTML(http.StatusOK, "private/closure", cc.res)
}