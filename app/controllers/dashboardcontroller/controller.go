package dashboardcontroller

import (
	"github.com/gin-gonic/gin"
)

type responseDashboard struct {
	Title   string
	Content string
	UserCurrent struct {
		UserName string
		Email    string
	}
}

type dashboardController struct {
	res *responseDashboard
}

type DashboardController interface {
	Dashboard(c *gin.Context)
}

func InitDashboardController() DashboardController {
	res := &responseDashboard{
		Title:   "La Belfortaine - Tableau de bord",
		Content: "dashboard",
	}

	return &dashboardController{res}
}
