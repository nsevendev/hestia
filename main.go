package main

import (
	"hestia/app/router"
	"html/template"

	"github.com/gin-gonic/gin"
)

func main() {
	serv := gin.Default()

	serv.SetFuncMap(template.FuncMap{})
	serv.LoadHTMLGlob("app/views/**/*.html")
	serv.Static("/assets", "./app/views/assets")

	router.Router(serv)

	serv.Run("0.0.0.0:4200")
}
