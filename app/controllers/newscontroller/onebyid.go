package newscontroller

import (
	"hestia/internal/logger"
	"hestia/internal/models"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (nc *newsController) OneById(c *gin.Context) {
	uuid := c.Params.ByName("uuid")

	listNews, errListNews := nc.newsService.GetAll(c)
	if errListNews != nil {
		logger.Errorf("[newscontroller::OneById] Erreur recuperation news : %v", errListNews)
		c.Redirect(
			http.StatusSeeOther, 
			"/dashboard/news?statusCode=" + strconv.Itoa(http.StatusBadRequest) + "&error=" + url.QueryEscape("Erreur lors de la récupération des actualités"),
		)
		return
	}

	oneNewsMap := make(map[string]*models.News)
	for i := range listNews {
		news := &listNews[i]
		oneNewsMap[news.UUID.String()] = news
	}

	news, found := oneNewsMap[uuid]
	if !found {
		logger.Errorf("[newscontroller::OneById] Erreur recuperation news: erreur à la récuperation de l'actualité")
		c.Redirect(
			http.StatusSeeOther, 
			"/dashboard/news?statusCode=" + strconv.Itoa(http.StatusBadRequest) + "&error=" + url.QueryEscape("Erreur lors de la récupération de l'actualité"),
		)
		return
	}

	nc.res.ListNews = listNews
	nc.res.News = news

	c.HTML(http.StatusOK, "private/news",nc. res)
}