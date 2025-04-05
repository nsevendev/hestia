package auth

import (
	"hestia/internal/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        userID := session.GetUserID(c)
        if userID == "" {
            c.Redirect(http.StatusSeeOther, "/login")
            c.Abort()
            return
        }
        c.Next()
    }
}