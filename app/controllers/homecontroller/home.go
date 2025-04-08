package homecontroller

import (
	"net/http"

	"hestia/internal/models"

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

	// Appeler GetAll en transmettant le contexte de la requête
	newsList, err := h.newsService.GetAll(c.Request.Context())
	if err != nil {
		h.res.ListNews = []models.News{} // On affecte une tranche vide en cas d'erreur
		h.res.Error = "Erreur lors de la récupération des actualités"
	} else {
		h.res.ListNews = newsList
	}

	c.HTML(http.StatusOK, "public/home", h.res)
}
