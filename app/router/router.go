package router

import (
	"hestia/app/controllers/dashboardcontroller"
	"hestia/app/controllers/gallerycontroller"
	"hestia/app/controllers/homecontroller"
	"hestia/app/controllers/newscontroller"
	"hestia/app/controllers/termscontroller"
	depinject "hestia/app/depInject"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine, container *depinject.Container) {

	// ╔═══════════════════════════════════════════════════════════╗
	// ║                 DECLARATION CONTROLLER                    ║
	// ╚═══════════════════════════════════════════════════════════╝

	home := homecontroller.InitHomeController()
	dash := dashboardcontroller.InitDashboardController()
	news := newscontroller.InitNewsController(container)
	gallery := gallerycontroller.InitGalleryController()
	terms := termscontroller.InitTermsController()

	// ╔═══════════════════════════════════════════════════════════╗
	// ║                        PARTIE SITE                        ║
	// ╚═══════════════════════════════════════════════════════════╝

	r.GET("/", home.Home)

	// ╔═══════════════════════════════════════════════════════════╗
	// ║                 PARTIE ADMIN DASHBOARD                    ║
	// ╚═══════════════════════════════════════════════════════════╝
	
	routeDashboard := r.Group("/dashboard")
	routeDashboard.GET("/", dash.Dashboard)

	// ╔═══════════════════════════════════════════════════════════╗
	// ║                     PARTIE ADMIN NEWS                     ║
	// ╚═══════════════════════════════════════════════════════════╝

	routeDashboard.GET("/news", news.List)
	routeDashboard.GET("/news/:uuid", news.OneById)
	routeDashboard.POST("/news", news.Create)
	routeDashboard.POST("/news/update/:uuid", news.UpdateById)
	routeDashboard.POST("/news/delete/:uuid", news.DeleteById)

	// ╔═══════════════════════════════════════════════════════════╗
	// ║                   PARTIE ADMIN GALLERY                    ║
	// ╚═══════════════════════════════════════════════════════════╝

	routeDashboard.GET("/gallery", gallery.First)

	// ╔═══════════════════════════════════════════════════════════╗
	// ║                     PARTIE ADMIN TERM                     ║
	// ╚═══════════════════════════════════════════════════════════╝

	routeDashboard.GET("/terms", terms.List)
}
