package homecontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *homeController) Home(c *gin.Context) {
	period, err := h.closurePeriodService.Active(c)
	if err != nil {
		h.res.PeriodClosed = nil
		h.res.Error = "Erreur lors de la récupération de la période de fermeture active"
		c.HTML(http.StatusInternalServerError, "public/home", h.res)
		return
	}

	h.res.PeriodClosed = period

	c.HTML(http.StatusOK, "public/home", h.res)
}