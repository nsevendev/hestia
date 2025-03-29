package router

import (
	"hestia/app/controllers"

	"github.com/gin-gonic/gin"
)

func Router(serv *gin.Engine) {
	// Routes de base
	serv.GET("/", controllers.Home)

	// Groupe de routes sous /dashboard
	dashboard := serv.Group("/dashboard")

	dashboard.GET("/", controllers.Dashboard)

	dashboard.GET("/news", controllers.GetAllNews)
	dashboard.POST("/news", controllers.CreateNews)
	dashboard.PUT("/news/:id", controllers.UpdateNews)
	dashboard.DELETE("/news/:id", controllers.DeleteNews)
}
