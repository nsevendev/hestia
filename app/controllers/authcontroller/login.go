package authcontroller

import (
	"hestia/internal/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ac *authController) Login(c *gin.Context) {
    email := c.PostForm("email")
    password := c.PostForm("password")

    user, err := ac.serviceAuth.Authenticate(c.Request.Context(), email, password)
    if err != nil {
        c.Redirect(http.StatusSeeOther, "/login?error=Email ou mot de passe invalide " + err.Error())
        return
    }

    session.SetUserSession(c, user.UUID.String())
    c.Redirect(http.StatusSeeOther, "/dashboard")
}