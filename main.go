package main

import (
	"hestia/app/router"
	_ "hestia/init"
	"hestia/internal/logger"
	"html/template"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func extractBacktickContent(s string) string {
	start := strings.Index(s, "`")
	end := strings.LastIndex(s, "`")

	if start == -1 || end == -1 || start == end {
		return "" // ou une erreur si tu préfères
	}

	return s[start+1 : end]
}

func main() {	
	serv := gin.Default()

	serv.SetFuncMap(template.FuncMap{})
	serv.LoadHTMLGlob("app/views/**/*.html")
	serv.Static("/assets", "./app/views/assets")

	router.Router(serv)

	port := os.Getenv("PORT")
	host := "0.0.0.0"
	hostTraefik := extractBacktickContent(os.Getenv("HOST_TRAEFIK"))

	logger.Success("Server is running on " + host + ":" + port)
	logger.Successf("Server is running on https://%v", hostTraefik)

	serv.Run(host + ":" + port)
}
 