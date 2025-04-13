package homecontroller

import (
	"hestia/internal/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *homeController) Home(c *gin.Context) {

	// Période de fermeture
	period, err := h.closurePeriodService.Active(c)
	if err != nil {
		logger.Warnf("[home] Erreur récupération période de fermeture : %v", err)
	}
	h.res.PeriodClosed = period

	// Actualités
	newsList, err := h.newsService.GetAll(c.Request.Context())
	if err != nil {
		logger.Warnf("[home] Erreur récupération des news : %v", err)
		h.res.ListNews = nil
	} else {
		h.res.ListNews = newsList
	}

	// Galerie (images)
	gallery, err := h.galleryService.GetFirst()
	if err != nil {
		logger.Warnf("[home] Erreur récupération galerie : %v", err)
		h.res.Gallery = nil
	} else {
		h.res.Gallery = gallery
	}

	// Render
	c.HTML(http.StatusOK, "public/home", h.res)
}
