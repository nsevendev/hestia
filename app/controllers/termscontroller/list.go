package termscontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (tc *termsController) List(c *gin.Context) {
	c.HTML(http.StatusOK, "private/terms", tc.res)
}
