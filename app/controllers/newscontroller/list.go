package newscontroller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (nc *newsController) List(c *gin.Context) {
	statusCodeStr := c.Query("statusCode")
	statusCode, _ := strconv.Atoi(statusCodeStr)

	listNews, err := nc.newsService.GetAll(c)
	if err != nil {
		nc.res.Error = "Erreur lors de la récupération des actualités"
		c.HTML(statusCode, "private/news", nc.res)
		return
	}

	nc.res.ListNews = listNews
	nc.res.News = nil
	nc.res.Success = c.Query("success")
	nc.res.Error = c.Query("error")

	c.HTML(http.StatusOK, "private/news", nc.res)
}
