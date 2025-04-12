package newscontroller

import (
	depinject "hestia/app/depInject"
	"hestia/internal/models"
	"hestia/internal/news"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ╔═══════════════════════════════════════════════════════════╗
// ║                         PRIVATE                   		   ║
// ╚═══════════════════════════════════════════════════════════╝

type responseNews struct {
	Title   string
	Content string
	ListNews []models.News
	News   *models.News
	Error   string
	Success string
}

type newsController struct {
	res *responseNews
	newsService news.NewsService
}

func validateDataStringEmpty(c *gin.Context , value *string, name string) string {
	if *value == "" {
		c.Redirect(
			http.StatusSeeOther, 
			"/dashboard/news?statusCode=" + strconv.Itoa(http.StatusBadRequest) + "&error=" + url.QueryEscape("Erreur de validation, " + name + " requis"),
		)
	}

	return  *value
}

// ╔═══════════════════════════════════════════════════════════╗
// ║                          PUBLIC	   	                   ║
// ╚═══════════════════════════════════════════════════════════╝

type NewsController interface {
	List(c *gin.Context)
	OneById(c *gin.Context)
	Create(c *gin.Context)
	UpdateById(c *gin.Context)
	DeleteById(c *gin.Context)
}

func InitNewsController(c *depinject.Container) NewsController {
	res := &responseNews{
		Title:   "news",
		Content: "news",
	}

	return &newsController{
		res, 
		c.NewsService,
	}
}

