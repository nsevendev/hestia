package main

import (
	"hestia/app/router"
	initializer "hestia/init"
	"hestia/internal/logger"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

/*
extraire le contenu entre les backticks

params:
	- s : string
return:
	- string : contenu entre les backticks
*/
func extractBacktickContent(s string) string {
	start := strings.Index(s, "`")
	end := strings.LastIndex(s, "`")

	if start == -1 || end == -1 || start == end {
		return ""
	}

	return s[start+1 : end]
}

/*
charge tout les templates html et les parse dans le moteur de template

params:
	- funcMap : FuncMap moteur de template
return:
	- template : template
*/
func loadTemplates(funcMap *template.FuncMap) *template.Template {
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

	return template.Must(template.New("").Funcs(*funcMap).ParseFiles(tmpl...))
}

func main() {

	// ╔═══════════════════════════════════════════════════════════╗
	// ║                       INIT SERVER                         ║
	// ╚═══════════════════════════════════════════════════════════╝

	initializer.Bootstrap()
	serv := gin.Default()
	hostTraefik := extractBacktickContent(os.Getenv("HOST_TRAEFIK"))

	// ╔═══════════════════════════════════════════════════════════╗
	// ║                       INIT TEMPLATE                       ║
	// ╚═══════════════════════════════════════════════════════════╝

	funcMap := template.FuncMap{
		"formatDate": func(t time.Time, layout string) string {
			if t.IsZero() {
				return "-- Aucune date --"
			}
			return t.Format(layout)
		},
		"hostTraefik": func() string {
        	return "https://" + hostTraefik
   		},
	}
	serv.SetHTMLTemplate(loadTemplates(&funcMap))
	serv.Static("/assets", "./app/views/assets")

	// ╔═══════════════════════════════════════════════════════════╗
	// ║                       INIT ROUTER                         ║
	// ╚═══════════════════════════════════════════════════════════╝

	router.Router(serv, initializer.Container)

	// ╔═══════════════════════════════════════════════════════════╗
	// ║           LOGGER OUT SERVER INFOS AFTER START             ║
	// ╚═══════════════════════════════════════════════════════════╝

	serv.NoRoute(func(c *gin.Context) {
		logger.Errorf("❌ Route inconnue : %s %s", c.Request.Method, c.Request.URL.Path)
		c.String(http.StatusNotFound, "404 not found")
	})

	port := os.Getenv("PORT")
	host := "0.0.0.0"
	logger.Success("Server is running on in container docker : " + host + ":" + port)
	logger.Successf("Server is running on navigator on : https://%v", hostTraefik)
	serv.Run(host + ":" + port)
}
