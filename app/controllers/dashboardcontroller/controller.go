package dashboardcontroller

import "github.com/gin-gonic/gin"

// ╔═══════════════════════════════════════════════════════════╗
// ║                            PRIVATE                        ║
// ╚═══════════════════════════════════════════════════════════╝

type responseDashboard struct {
	Title   string
	Content string
}

type dashboardController struct {
	res *responseDashboard
}

// ╔═══════════════════════════════════════════════════════════╗
// ║                          PUBLIC	   	                   ║
// ╚═══════════════════════════════════════════════════════════╝

type DashboardController interface {
	Dashboard(c *gin.Context)
}

func InitDashboardController() DashboardController {
	res := &responseDashboard{
		Title:   "Dashboard",
		Content: "dashboard",
	}

	return &dashboardController{res}
}