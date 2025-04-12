package authcontroller

import (
	"hestia/internal/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ac *authController) ShowLogin(c *gin.Context) {
	ac.res.Error = c.Query("error")
	userId := session.GetUserID(c)
	if userId != "" {
		c.Redirect(http.StatusSeeOther, "/dashboard")
	}
    c.HTML(http.StatusOK, "public/login", ac.res)
}