package main

import (
	"hestia/app/router"
	_ "hestia/init"
	"hestia/internal/logger"
	"html/template"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {	
	serv := gin.Default()

	serv.SetFuncMap(template.FuncMap{})
	serv.LoadHTMLGlob("app/views/**/*.html")
	serv.Static("/assets", "./app/views/assets")

	router.Router(serv)

	port := os.Getenv("PORT")
	host := "0.0.0.0"
	hostTraefik := os.Getenv("HOST_TRAEFIK")

	logger.Success("Server is running on " + host + ":" + port)
	logger.Successf("Server is running on https://%v", hostTraefik)

	serv.Run(host + ":" + port)
}
 