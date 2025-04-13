package dashboardcontroller

import (
	"hestia/internal/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (d *dashboardController) Dashboard(c *gin.Context) {
	userId := session.GetUserID(c)
	if userId != "" {
		email, userName := session.GetUserInfos(c)
		d.res.UserCurrent.UserName = userName
		d.res.UserCurrent.Email = email
		c.HTML(http.StatusOK, "private/dashboard", d.res)
	}
	c.Redirect(http.StatusSeeOther, "/login")
}