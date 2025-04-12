package closureperiodcontroller

import (
	depinject "hestia/app/depInject"
	"hestia/internal/closedperiod"
	"hestia/internal/models"

	"github.com/gin-gonic/gin"
)

// ╔═══════════════════════════════════════════════════════════╗
// ║                            PRIVATE                        ║
// ╚═══════════════════════════════════════════════════════════╝

type responseClosurePeriod struct {
	Title   string
	Content string
	ListPeriod []models.ClosurePeriod
	Error string
	Success string
}

type closurePeriodController struct {
	res *responseClosurePeriod
	closureService closedperiod.ClosedPeriodService
}


// ╔═══════════════════════════════════════════════════════════╗
// ║                            PUBLIC                         ║
// ╚═══════════════════════════════════════════════════════════╝

type ClosurePeriodController interface {
	List(c *gin.Context)
	Create(c *gin.Context)
	DeleteById(c *gin.Context)
}

func InitHomeController(c *depinject.Container) ClosurePeriodController {
	res := &responseClosurePeriod{
		Title:  "Periodes de fermeture",
		Content: "closure",
	}

	return &closurePeriodController{res, c.ClosedPeriodService}
}