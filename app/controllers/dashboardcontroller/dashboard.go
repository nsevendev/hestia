package dashboardcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (d *dashboardController) Dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "private/dashboard", d.res)
}