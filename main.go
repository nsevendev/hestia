package main

import (
	"hestia/app/router"
	_ "hestia/init"
	"hestia/internal/logger"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func extractBacktickContent(s string) string {
	start := strings.Index(s, "`")
	end := strings.LastIndex(s, "`")

	if start == -1 || end == -1 || start == end {
		return ""
	}

	return s[start+1 : end]
}

func loadTemplates() *template.Template {
	var tmpl []string
	
	partials, _ := filepath.Glob("app/views/partials/*.html")
	layout, _ := filepath.Glob("app/views/layouts/*.html")
	pagesPublic, _ := filepath.Glob("app/views/public/*.html")
	pagesPublicChild, _ := filepath.Glob("app/views/public/**/*.html")
	pagesPrivate, _ := filepath.Glob("app/views/private/*.html")
	pagesPrivateChild, _ := filepath.Glob("app/views/private/**/*.html")

	tmpl = append(tmpl, partials...)
	tmpl = append(tmpl, layout...)
	tmpl = append(tmpl, pagesPublic...)
	tmpl = append(tmpl, pagesPublicChild...)
	tmpl = append(tmpl, pagesPrivate...)
	tmpl = append(tmpl, pagesPrivateChild...)

	logger.Infof("Templates chargés: %v", tmpl)

	return template.Must(template.ParseFiles(tmpl...))
}

func main() {	
	serv := gin.Default()
	
	//serv.Delims("{[{", "}]}") // ajouter les délimiteurs pour le moteur de template utilisé pour des moldels (pipe) perso
	
	serv.SetFuncMap(template.FuncMap{}) // ajouter justement des models custom ici doc : https://gin-gonic.com/docs/examples/html-rendering/
	
	serv.SetHTMLTemplate(loadTemplates())
	
	serv.Static("/assets", "./app/views/assets")

	router.Router(serv)

	port := os.Getenv("PORT") // que pour du log
	host := "0.0.0.0" // que pour du log
	hostTraefik := extractBacktickContent(os.Getenv("HOST_TRAEFIK")) // que pour du log

	logger.Success("Server is running on " + host + ":" + port)
	logger.Successf("Server is running on https://%v", hostTraefik)

	serv.Run(host + ":" + port)
}
 