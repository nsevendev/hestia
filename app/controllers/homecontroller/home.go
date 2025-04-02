package homecontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *homeController) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "public/home", h.res)
}