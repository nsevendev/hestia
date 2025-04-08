package homecontroller

import (
	depinject "hestia/app/depInject"
	"hestia/internal/closedperiod"
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
	Error        string
}

type homeController struct {
	res                  *responseHome
	closurePeriodService closedperiod.ClosedPeriodService
	newsService          news.NewsService
}

// ╔═══════════════════════════════════════════════════════════╗
// ║                            PUBLIC                         ║
// ╚═══════════════════════════════════════════════════════════╝

type HomeController interface {
	Home(c *gin.Context)
}

// Initialisation du contrôleur en injectant les services requis.
func InitHomeController(c *depinject.Container) HomeController {
	res := &responseHome{
		Title:   "La Belfortaine - Boucherie & Charcuterie traditionnelle à Belfort",
		Content: "home",
	}

	return &homeController{
		res:                  res,
		closurePeriodService: c.ClosedPeriodService,
		newsService:          c.NewsService, // Assurez-vous que c.NewsService est correctement défini dans votre container.
	}
}
