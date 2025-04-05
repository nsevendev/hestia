package authcontroller

import (
	"hestia/internal/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ac *authController) Logout(c *gin.Context) {
    session.ClearSession(c)
    c.Redirect(http.StatusSeeOther, "/login")
}