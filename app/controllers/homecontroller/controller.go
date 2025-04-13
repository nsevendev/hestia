package homecontroller

import (
	depinject "hestia/app/depInject"
	"hestia/internal/closedperiod"
	"hestia/internal/gallery"
	"hestia/internal/models"
	"hestia/internal/news"

	"github.com/gin-gonic/gin"
)

// ╔═══════════════════════════════════════════════════════════╗
// ║                            PRIVATE                        ║
// ╚═══════════════════════════════════════════════════════════╝

type responseHome struct {
	Title        string
	Content      string
	PeriodClosed *models.ClosurePeriod
	ListNews     []models.News
	Gallery      *models.Gallery
	Error        string
}

type homeController struct {
	res                  *responseHome
	closurePeriodService closedperiod.ClosedPeriodService
	newsService          news.NewsService
	galleryService       gallery.GalleryService
}

// ╔═══════════════════════════════════════════════════════════╗
// ║                            PUBLIC                         ║
// ╚═══════════════════════════════════════════════════════════╝

type HomeController interface {
	Home(c *gin.Context)
}

func InitHomeController(c *depinject.Container) HomeController {
	res := &responseHome{
		Title:   "La Belfortaine - Boucherie & Charcuterie traditionnelle à Belfort",
		Content: "home",
	}

	return &homeController{
		res:                  res,
		closurePeriodService: c.ClosedPeriodService,
		newsService:          c.NewsService,
		galleryService:       c.GalleryService,
	}
}
