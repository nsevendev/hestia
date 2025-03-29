package router

import (
	"hestia/app/controllers"

	"github.com/gin-gonic/gin"
)

func Router(serv *gin.Engine) {
	serv.GET("/", controllers.Home)
	serv.GET("/dashboard", controllers.Test)
}