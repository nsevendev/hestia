package router

import (
	"hestia/app/controllers/authcontroller"
	"hestia/app/controllers/dashboardcontroller"
	"hestia/app/controllers/gallerycontroller"
	"hestia/app/controllers/homecontroller"
	"hestia/app/controllers/newscontroller"
	"hestia/app/controllers/termscontroller"
	depinject "hestia/app/depInject"
	"hestia/internal/auth"
	"hestia/internal/session"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine, container *depinject.Container) {
	r.Use(session.Init("mykey"))

	// ╔═══════════════════════════════════════════════════════════╗
	// ║                 DECLARATION CONTROLLER                    ║
	// ╚═══════════════════════════════════════════════════════════╝

	authen := authcontroller.InitHomeController(container)
	home := homecontroller.InitHomeController()
	dash := dashboardcontroller.InitDashboardController()
	news := newscontroller.InitNewsController(container)
	gallery := gallerycontroller.InitGalleryController(container)
	terms := termscontroller.InitTermsController()

	// ╔═══════════════════════════════════════════════════════════╗
	// ║                        PARTIE SITE                        ║
	// ╚═══════════════════════════════════════════════════════════╝

	r.GET("/", home.Home)

	// ╔═══════════════════════════════════════════════════════════╗
	// ║                        PARTIE AUTH                        ║
	// ╚═══════════════════════════════════════════════════════════╝

	r.GET("/login", authen.ShowLogin)
	r.POST("/login", authen.Login)
	r.GET("/logout", authen.Logout)

	// ╔═══════════════════════════════════════════════════════════╗
	// ║                 PARTIE ADMIN DASHBOARD                    ║
	// ╚═══════════════════════════════════════════════════════════╝
	
	routeDashboard := r.Group("/dashboard")
	routeDashboard.Use(auth.RequireAuth())

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
	routeDashboard.POST("/gallery", gallery.AddImage)
	routeDashboard.POST("/gallery/delete/:uuid", gallery.DeleteImageById)

	// ╔═══════════════════════════════════════════════════════════╗
	// ║                     PARTIE ADMIN TERM                     ║
	// ╚═══════════════════════════════════════════════════════════╝

	routeDashboard.GET("/terms", terms.List)
}
