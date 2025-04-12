package homecontroller

import (
	depinject "hestia/app/depInject"
	"hestia/internal/closedperiod"
	"hestia/internal/models"

	"github.com/gin-gonic/gin"
)

// ╔═══════════════════════════════════════════════════════════╗
// ║                            PRIVATE                        ║
// ╚═══════════════════════════════════════════════════════════╝

type responseHome struct {
	Title   string
	Content string
	PeriodClosed *models.ClosurePeriod
	Error string
}

type homeController struct {
	res *responseHome
	closurePeriodService closedperiod.ClosedPeriodService
}


// ╔═══════════════════════════════════════════════════════════╗
// ║                            PUBLIC                         ║
// ╚═══════════════════════════════════════════════════════════╝

type HomeController interface {
	Home(c *gin.Context)
}

func InitHomeController(c *depinject.Container) HomeController {
	res := &responseHome{
		Title:  "La Belfortaine - Boucherie & Charcuterie traditionnelle à Belfort",
		Content: "home",
	}

	return &homeController{res, c.ClosedPeriodService}
}