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
	dashboard.GET("/news/:uuid", controllers.GetOneNews)
	dashboard.POST("/news", controllers.CreateNews)
	dashboard.POST("/news/update/:uuid", controllers.UpdateNews)
	dashboard.POST("/news/delete/:uuid", controllers.DeleteNews)
}
